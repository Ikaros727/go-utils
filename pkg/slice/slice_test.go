package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_Exist(t *testing.T) {
	ast := assert.New(t)
	strList := []string{"a", "b", "c"}

	// determines whether the target is in the s
	ast.Equal(Exist(strList, "a"), true, "slice.Exist must return true value")
	ast.Equal(Exist(strList, "d"), false, "slice.Exist must return false value")
	// determines whether the target is not in the s
	ast.Equal(!Exist(strList, "a"), false, "slice.Exist must return false value")
	ast.Equal(!Exist(strList, "d"), true, "slice.Exist must return true value")
}

func TestSlice_Distinct(t *testing.T) {
	ast := assert.New(t)
	strList := []string{"b", "a", "b", "c", "c", "a"}

	distinct := Distinct(strList)
	ast.Equal("b", distinct[0])
	ast.Equal("a", distinct[1])
	ast.Equal("c", distinct[2])
}

func TestSlice_Union(t *testing.T) {
	ast := assert.New(t)
	sliceA := []string{"a", "b"}
	sliceB := []string{"c", "d"}

	union := Union(sliceA, sliceB)
	ast.Equal(4, len(union))
	if len(union) == 4 {
		ast.Equal("a", union[0])
		ast.Equal("b", union[1])
		ast.Equal("c", union[2])
		ast.Equal("d", union[3])
	}
}

func TestSlice_Intersect(t *testing.T) {
	ast := assert.New(t)
	sliceA := []string{"a", "b"}
	sliceB := []string{"b", "c"}

	intersect := Intersect(sliceA, sliceB)
	ast.Equal(1, len(intersect))
	if len(intersect) == 1 {
		ast.Equal("b", intersect[0])
	}
}

func TestSlice_Subtract(t *testing.T) {
	ast := assert.New(t)
	sliceA := []string{"a", "b"}
	sliceB := []string{"b", "c"}

	subtract := Subtract(sliceA, sliceB)
	ast.Equal(1, len(subtract))
	if len(subtract) == 1 {
		ast.Equal("a", subtract[0])
	}
}

func TestSlice_ExclusiveOr(t *testing.T) {
	ast := assert.New(t)
	sliceA := []string{"a", "b"}
	sliceB := []string{"b", "c"}

	subtract := ExclusiveOr(sliceA, sliceB)
	ast.Equal(2, len(subtract))
	if len(subtract) == 2 {
		ast.Equal("a", subtract[0])
		ast.Equal("c", subtract[1])
	}
}
