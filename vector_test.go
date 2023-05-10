package toolkit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type VectorTest struct {
	a string
}

func Test_InVectors(t *testing.T) {
	r := InVectors[string]("ss1", Vector[string]{"sss", "ss1"})
	assert.Equal(t, true, r)
}

func Test_Vector_Contains(t *testing.T) {
	r := Vector[string]{"sss", "ssccc"}.Contains("sss")
	assert.Equal(t, true, r)

	r = Vector[*VectorTest]{&VectorTest{}, &VectorTest{}}.Contains(&VectorTest{})
	assert.Equal(t, false, r)
	a := &VectorTest{}
	r = Vector[*VectorTest]{&VectorTest{}, a}.Contains(a)
	assert.Equal(t, true, r)
}
