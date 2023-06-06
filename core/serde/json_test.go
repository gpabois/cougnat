package serde

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Json(t *testing.T) {
	expectedVal := fixture()

	resMarshal := MarshalJson(expectedVal)
	assert.True(t, resMarshal.IsSuccess(), resMarshal.UnwrapError())

	marshalled := resMarshal.Expect()
	fmt.Println(marshalled)

	resUnMarshal := UnMarshalJson[testStruct](marshalled)
	assert.True(t, resUnMarshal.IsSuccess(), resUnMarshal.UnwrapError())

	unmarshalled := resUnMarshal.Expect()
	assert.Equal(t, expectedVal, unmarshalled)
}
