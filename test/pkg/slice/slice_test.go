package slice

import (
	"testing"

	"github.com/Icarus-0727/go-utils/pkg/slice"
	"github.com/stretchr/testify/assert"
)

func TestSlice_Exist(t *testing.T) {
	ast := assert.New(t)

	strList := []string{"a", "b", "c"}

	ast.Equal(slice.Exist(strList, "a"), true, "slice.Exist must return true value")
	ast.Equal(slice.Exist(strList, "d"), false, "slice.Exist must return false value")
}

func TestSlice_NotExist(t *testing.T) {
	ast := assert.New(t)

	strList := []string{"a", "b", "c"}

	ast.Equal(slice.NotExist(strList, "a"), false, "slice.Exist must return false value")
	ast.Equal(slice.NotExist(strList, "d"), true, "slice.Exist must return true value")
}
