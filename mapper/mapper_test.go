package mapper

import (
	"testing"
)

func TestMapper(t *testing.T) {

	mapp := MapperRecord{}
	if len(mapp.Fields) != 0 {
		t.Error("Expected empty mapper, got", len(mapp.Fields))
	}
}
