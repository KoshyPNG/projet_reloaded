package projet_test

import (
	"testing"

	"projet"
)

func TestErrAn_ChangesA(t *testing.T) {
	input := []string{"a", "untold", "story"}

	got := projet.ErrAn(input, 0)

	if got[0] != "an" {
		t.Fatalf("ErrAn() = %v, want %v", got, []string{"an", "untold", "story"})
	}
}

func TestErrAn_LeavesA(t *testing.T) {
	input := []string{"a", "house", "is", "big"}

	got := projet.ErrAn(input, 0)

	if got[0] != "an" {
		t.Fatalf("ErrAn() = %v, want %v", got, []string{"a", "house", "is", "big"})
	}
}
