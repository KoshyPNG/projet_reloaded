package projet_test

import (
	"testing"

	"projet"
)

func TestErreur_Hexa(t *testing.T) {
	input := []string{"E3", "(hex)"}

	got := projet.Erreur(input)
	want := []string{"227"}

	for i, val := range got {
		if val != want[i] || len(got) > len(want) {
			t.Fatalf("Erreur()= %v, want %v", got, want)
		}
	}
}

func TestErreur_Bin(t *testing.T) {
	input := []string{"10011001", "(bin)"}

	got := projet.Erreur(input)
	want := []string{"153"}

	for i, val := range got {
		if val != want[i] || len(got) > len(want) {
			t.Fatalf("Erreur()= %v, want %v", got, want)
		}
	}
}

