package projet_test

import (
	"reflect"
	"testing"

	"projet"
)

func TestFormat_RemovesEmptyStrings(t *testing.T) {
	input := []string{"hello", "", "world"}

	got := projet.Format(input)
	want := []string{"hello", "world"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Format() = %v, want %v", got, want)
	}
}

func TestErreur_Up(t *testing.T) {
	input := []string{"hello", "(up)"}

	got := projet.Erreur(input)
	want := []string{"HELLO"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Erreur() = %v, want %v", got, want)
	}
}

func TestPonctuation_MovesComma(t *testing.T) {
	input := []string{"hello", ",", "world"}

	got := projet.Ponctuation(input)
	want := []string{"hello,", "world"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Ponctuation() = %v, want %v", got, want)
	}
}
