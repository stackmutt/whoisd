package mapper

import (
	"testing"
)

func TestMapper(t *testing.T) {

	mapp := MapperRecord{}
	if mapp.Count() != 0 {
		t.Error("Expected empty mapper record, got", mapp.Count())
	}
}
