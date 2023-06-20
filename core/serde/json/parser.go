package json

import (
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/gpabois/cougnat/core/result"
	"golang.org/x/exp/slices"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		n   int
	}
}

var UnexpectedEof error = errors.New("unexpected eof")

func UnexpectedToken(token Token) error {
	return errors.New("unexpected token")
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Skip useless tokens (such as whitespaces)
func (p *Parser) skip(tokTypes ...TokenType) {
	for {
		if !slices.Contains(tokTypes, p.scan().typ) {
			p.rewind()
			return
		}
	}
}

func (p *Parser) Parse() result.Result[Json] {
	p.skip(WS)
	return p.parseRoot()
}

// Parse the root of the Json
func (p *Parser) parseRoot() result.Result[Json] {
	tok := p.scan()
	if tok.typ == LEFT_ARRAY {
		p.rewind()
		arrRes := p.parseArray()

		if arrRes.HasFailed() {
			return result.Result[Json]{}.Failed(arrRes.UnwrapError())
		}

		return result.Success(Json{}.Array(arrRes.Expect()))
	} else if tok.typ == LEFT_DOCUMENT {
		p.rewind()
		docRes := p.parseDocument()
		if docRes.HasFailed() {
			return result.Result[Json]{}.Failed(docRes.UnwrapError())
		}
		return result.Success(Json{}.Document(docRes.Expect()))
	} else {
		return result.Result[Json]{}.Failed(UnexpectedToken(tok))
	}
}

// Parse an array of values
func (p *Parser) parseArray() result.Result[Array] {
	p.scan() // Scan the [ out
	var values []Value

	for {
		p.skip(WS, COMMA)
		tok := p.scan()
		if tok.typ == EOF {
			return result.Result[Array]{}.Failed(UnexpectedEof)
		} else if tok.typ != RIGHT_ARRAY {
			p.rewind()
			valRes := p.parseValue()

			if valRes.HasFailed() {
				return result.Result[Array]{}.Failed(valRes.UnwrapError())
			}

			values = append(values, valRes.Expect())
		} else {
			return result.Success(Array{
				Elements: values,
			})
		}
	}
}

// Parse a document
func (p *Parser) parseDocument() result.Result[Document] {
	p.scan() // Remove the {
	var pairs []Element

	for {
		p.skip(WS, COMMA)

		tok := p.scan()
		// Reach the end of the document
		if tok.typ == RIGHT_DOCUMENT {
			return result.Success(Document{
				Pairs: pairs,
			})
		} else if tok.typ == EOF {
			return result.Result[Document]{}.Failed(UnexpectedEof)
		}

		p.rewind()
		pairRes := p.parseDocumentPair()
		if pairRes.HasFailed() {
			return result.Result[Document]{}.Failed(p.parseDocumentPair().UnwrapError())
		}
		pairs = append(pairs, pairRes.Expect())
	}
}

// A document-pair is a STRING (WS*) COLON (WS*) VALUE sequence
func (p *Parser) parseDocumentPair() result.Result[Element] {
	key := p.scan()

	if key.typ != STRING {
		return result.Result[Element]{}.Failed(UnexpectedToken(key))
	}

	p.skip(WS)
	if tok := p.scan(); tok.typ != COLON {
		return result.Result[Element]{}.Failed(UnexpectedToken(tok))
	}

	valueRes := p.parseValue()

	if valueRes.HasFailed() {
		return result.Result[Element]{}.Failed(valueRes.UnwrapError())
	}

	return result.Success(Element{
		Key:   key.lit,
		Value: valueRes.Expect(),
	})
}

// Parse a value
func (p *Parser) parseValue() result.Result[Value] {
	p.skip(WS)

	tok := p.scan()

	switch tok.typ {
	case LEFT_DOCUMENT:
		p.rewind()
		docRes := p.parseDocument()
		if docRes.HasFailed() {
			return result.Result[Value]{}.Failed(docRes.UnwrapError())
		}
		return result.Success(Value{}.Document(docRes.Expect()))
	case LEFT_ARRAY:
		p.rewind()
		arrRes := p.parseArray()
		if arrRes.HasFailed() {
			return result.Result[Value]{}.Failed(arrRes.UnwrapError())
		}
		return result.Success(Value{}.Array(arrRes.Expect()))
	case TRUE:
		return result.Success(Value{}.Bool(true))
	case FALSE:
		return result.Success(Value{}.Bool(false))
	case NULL:
		return result.Success(Value{}.Null())
	case STRING:
		return result.Success(Value{}.String(tok.lit))
	case NUMBER:
		// Treat it as float
		if strings.Contains(tok.lit, ".") {
			fval, err := strconv.ParseFloat(tok.lit, 64)
			if err != nil {
				return result.Result[Value]{}.Failed(err)
			}
			return result.Success(Value{}.Float(fval))
		} else {
			ival, err := strconv.ParseInt(tok.lit, 10, 64)
			if err != nil {
				return result.Result[Value]{}.Failed(err)
			}
			return result.Success(Value{}.Integer(int(ival)))
		}
	case EOF:
		return result.Result[Value]{}.Failed(UnexpectedEof)
	default:
		return result.Result[Value]{}.Failed(UnexpectedToken(tok))
	}
}

// Scan the next token
func (p *Parser) scan() Token {
	if p.buf.n > 0 {
		p.buf.n = 0
		return p.buf.tok
	}

	p.buf.tok = p.s.Scan()
	return p.buf.tok
}

// Rewind the scanner by 1
func (p *Parser) rewind() {
	p.buf.n = 1
}
