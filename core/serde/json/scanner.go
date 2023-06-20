package json

import (
	"bufio"
	"bytes"
	"io"

	"github.com/gpabois/cougnat/core/ops"
)

type TokenType int
type Token struct {
	typ TokenType
	lit string
}

const (
	TOK_EOF = iota
	TOK_INVALID
	TOK_WS

	TOK_OPEN_DOCUMENT
	TOK_CLOSE_DOCUMENT
	TOK_COMMA
	TOK_COLON

	TOK_STRING
	TOK_TRUE
	TOK_FALSE
	TOK_NULL

	TOK_NUMBER

	TOK_OPEN_ARRAY
	TOK_CLOSE_ARRAY
)

const eof = rune(0)

func isEscape(ch rune) bool {
	return ch == '\\'
}

func isLetter(ch rune) bool {
	return ops.Within(ch, 'a', 'z') || ops.Within(ch, 'A', 'Z')
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isDigit(ch rune) bool {
	return ops.Within(ch, '0', '9')
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) rewind() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) scanIdent() Token {
	var buf bytes.Buffer

	for {
		ch := s.read()

		if !isLetter(ch) {
			s.rewind()
			break
		}

		buf.WriteRune(ch)
	}

	switch buf.String() {
	case "true":
		return Token{
			typ: TOK_TRUE,
			lit: buf.String(),
		}
	case "false":
		return Token{
			typ: TOK_FALSE,
			lit: buf.String(),
		}
	case "null":
		return Token{
			typ: TOK_NULL,
			lit: buf.String(),
		}
	default:
		return Token{
			typ: TOK_INVALID,
			lit: buf.String(),
		}
	}
}

func (s *Scanner) scanNumber() Token {
	var buf bytes.Buffer
	isFraction := false

	for {
		ch := s.read()
		if isDigit(ch) {
			buf.WriteRune(ch)
		} else if ch == '.' && !isFraction {
			isFraction = true
			buf.WriteRune(ch)
		} else {
			return Token{
				typ: TOK_INVALID,
				lit: "",
			}
		}
	}
}

func (s *Scanner) scanWhiteSpaces() Token {
	var buf bytes.Buffer
	for {
		ch := s.read()
		if !isWhiteSpace(ch) {
			s.rewind()
			return Token{
				typ: TOK_WS,
				lit: buf.String(),
			}
		}

		buf.WriteRune(ch)
	}
}
func (s *Scanner) scanString() Token {
	var buf bytes.Buffer
	prev := rune(0)
	for {
		ch := s.read()
		// We escaped the " character
		if ch == '"' && isEscape(prev) {
			buf.WriteRune(ch)
		} else if ch == '"' || ch == eof {
			s.rewind()
			return Token{
				typ: TOK_STRING,
				lit: buf.String(),
			}
		} else {
			buf.WriteRune(ch)
		}
		prev = ch
	}

}

func (s *Scanner) Scan() Token {
	// Read character
	ch := s.read()

	if ch == '"' {
		return s.scanString()
	} else if isWhiteSpace(ch) {
		return s.scanWhiteSpaces()
	} else if isDigit(ch) {
		s.rewind()
		return s.scanNumber()
	} else if isLetter(ch) {
		s.rewind()
		return s.scanIdent()
	} else if ch == ':' {
		return Token{
			typ: TOK_COLON,
			lit: ":",
		}
	} else if ch == ',' {
		return Token{
			typ: TOK_COMMA,
			lit: ",",
		}
	} else if ch == eof {
		return Token{
			typ: TOK_EOF,
			lit: "",
		}
	} else {
		return Token{
			typ: TOK_INVALID,
			lit: "",
		}
	}
}
