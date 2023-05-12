package result

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapResult(t *testing.T) {
	res := Map(Success(10), func(val int) bool { return val > 9 })

	assert.True(t, res.Expect())
}
