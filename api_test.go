package urbandictionary

import (
	"testing"
)

func TestQuery(t *testing.T) {
	r, err := Query("blumkin")
	if err != nil {
		t.Error(err)
	}
	if r.Type != "exact" {
		t.Errorf("Expected: exact, Obtained: %s", r.Type)
	}
}
