package serde

import (
	"fmt"
	"testing"

	"github.com/gpabois/cougnat/core/option"
	"github.com/stretchr/testify/assert"
)

func Test_Bson(t *testing.T) {
	expectedVal := testStruct{
		OptValue: option.Some("test"),
		StructValue: subTestStruct{
			El0: 10,
			El1: true,
		},
	}

	resMarshal := MarshalBson(expectedVal)
	assert.True(t, resMarshal.IsSuccess(), resMarshal.UnwrapError())

	marshalled := resMarshal.Expect()
	fmt.Println(marshalled)
	
	resUnMarshal := UnMarshalBson[testStruct](marshalled)
	assert.True(t, resUnMarshal.IsSuccess(), resUnMarshal.UnwrapError())

	unmarshalled := resUnMarshal.Expect()
	assert.Equal(t, expectedVal, unmarshalled)
}
