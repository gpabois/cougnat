package serde

import (
	"bufio"
	"bytes"

	"github.com/gpabois/cougnat/core/result"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

// Marshal a BSON Document
func MarshalBson[T any](value T) result.Result[[]byte] {
	var buffer bytes.Buffer
	bufferWriter := bufio.NewWriter(&buffer)
	// Create the stream writer
	writer, err := bsonrw.NewBSONValueWriter(bufferWriter)
	if err != nil {
		return result.Failed[[]byte](err)
	}

	// Create the encoder
	encoder, err := bson.NewEncoder(writer)
	if err != nil {
		return result.Failed[[]byte](err)
	}

	// Encode the normalised value
	normRes := Normalise(value)
	err = encoder.Encode(normRes)
	if err != nil {
		return result.Failed[[]byte](err)
	}

	bufferWriter.Flush()
	return result.Success(buffer.Bytes())
}

func UnMarshalBson[T any](raw []byte) result.Result[T] {
	var norm map[string]any
	reader := bsonrw.NewBSONDocumentReader(raw)

	decoder, err := bson.NewDecoder(reader)
	if err != nil {
		return result.Failed[T](err)
	}

	err = decoder.Decode(&norm)
	if err != nil {
		return result.Failed[T](err)
	}

	return DeNormalise[T](norm)
}
