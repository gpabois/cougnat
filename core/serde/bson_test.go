package serde

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bson(t *testing.T) {
	expectedVal := fixture()

	resMarshal := MarshalBson(expectedVal)
	assert.True(t, resMarshal.IsSuccess(), resMarshal.UnwrapError())

	marshalled := resMarshal.Expect()
	fmt.Println(marshalled)

	resUnMarshal := UnMarshalBson[testStruct](marshalled)
	assert.True(t, resUnMarshal.IsSuccess(), resUnMarshal.UnwrapError())

	unmarshalled := resUnMarshal.Expect()
	assert.Equal(t, expectedVal, unmarshalled)
}
