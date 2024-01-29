package slice

import (
	"testing"

	"github.com/Icarus-0727/go-utils/pkg/slice"
	"github.com/stretchr/testify/assert"
)

func TestSlice_Exist(t *testing.T) {
	ast := assert.New(t)
	strList := []string{"a", "b", "c"}

	// determines whether the target is in the s
	ast.Equal(slice.Exist(strList, "a"), true, "slice.Exist must return true value")
	ast.Equal(slice.Exist(strList, "d"), false, "slice.Exist must return false value")
	// determines whether the target is not in the s
	ast.Equal(!slice.Exist(strList, "a"), false, "slice.Exist must return false value")
	ast.Equal(!slice.Exist(strList, "d"), true, "slice.Exist must return true value")
}

func TestSlice_CustomizeExist(t *testing.T) {
	ast := assert.New(t)
	type TestModel struct{ Latter string }
	strList := []TestModel{{Latter: "a"}, {Latter: "b"}, {Latter: "c"}}

	// determines whether the target is in the s
	ast.Equal(slice.CustomizeExist(strList, TestModel{Latter: "a"},
		func(s, target TestModel) bool { return s.Latter == target.Latter }), true, "slice.Exist must return true value")
	ast.Equal(slice.CustomizeExist(strList, TestModel{Latter: "d"},
		func(s, target TestModel) bool { return s.Latter == target.Latter }), false, "slice.Exist must return false value")
	// determines whether the target is not in the s
	ast.Equal(!slice.CustomizeExist(strList, TestModel{Latter: "a"},
		func(s, target TestModel) bool { return s.Latter == target.Latter }), false, "slice.Exist must return false value")
	ast.Equal(!slice.CustomizeExist(strList, TestModel{Latter: "d"},
		func(s, target TestModel) bool { return s.Latter == target.Latter }), true, "slice.Exist must return true value")
}
