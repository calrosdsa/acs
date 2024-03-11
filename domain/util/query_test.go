package util_test

import (
	_query "acs/domain/util"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestArrayToString(t *testing.T) {
	t.Run("TestArrayString", func(t *testing.T) {
		res := _query.ArrayToString([]int{1, 2, 3, 4}, ",")
		assert.NotEmpty(t, res)
		assert.Equal(t, "1,2,3,4", res)
	})
}
