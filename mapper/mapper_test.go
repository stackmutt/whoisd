package mapper

import (
	"testing"
)

func TestMapper(t *testing.T) {

	bundle := make(Bundle, 2)
	bundle[0].TLDs = []string{"com", "net"}
	bundle[1].TLDs = []string{"tld"}
	if bundle.EntryByTLD("net").TLDs[0] != "com" {
		t.Error("Expected com, net entry, got", bundle.EntryByTLD("net").TLDs)
	}
	if bundle.EntryByTLD("tld").TLDs[0] != "tld" {
		t.Error("Expected tld entry, got", bundle.EntryByTLD("tld").TLDs)
	}
}
