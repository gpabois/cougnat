package serde

import (
	"testing"

	"github.com/gpabois/cougnat/core/option"
	"github.com/stretchr/testify/assert"
)

type subTestStruct struct {
	El0 int
	El1 []bool
}
type testStruct struct {
	OptValue    option.Option[string] `serde:"opt_value"`
	StructValue subTestStruct         `serde:"struct_value"`
}

func fixture() testStruct {
	return testStruct{
		OptValue: option.Some("test"),
		StructValue: subTestStruct{
			El0: 10,
			El1: []bool{true, false, true},
		},
	}
}

func Test_Normalisation(t *testing.T) {
	expectedVal := fixture()

	expectedNorm := make(NormalisedStruct)
	expectedNorm["opt_value"] = "test"
	expectedNorm["struct_value"] = NormalisedStruct{
		"El0": 10,
		"El1": []any{true, false, true},
	}

	norm := Normalise(expectedVal)
	assert.Equal(t, expectedNorm, norm)

	res := DeNormalise[testStruct](norm)
	assert.True(t, res.IsSuccess(), res.UnwrapError())

	val := res.Expect()
	assert.Equal(t, expectedVal, val)
}
