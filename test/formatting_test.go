package projet_test

import (
	"testing"

	"projet"
)

func TestFormat_RemovesEmptyStrings(t *testing.T) {
	input := []string{"hello", "", "world"}

	got := projet.Format(input)
	want := []string{"hello", "world"} 

	for i, val := range got {
		if val != want[i] || len(got) > len(want) {
			t.Fatalf("Erreur()= %v, want %v", got, want)
		}
	}
}

func TestReecrire(t *testing.T) {
	input := []string{"hello", "(up)"}

	got := projet.Reecrire(input,1)
	want := []string{"hello"}

	for i, val := range got {
		if val != want[i] || len(got) > len(want) {
			t.Fatalf("Erreur()= %v, want %v", got, want)
		}
	}
}

func TestErreur_Up(t *testing.T) {
	input := []string{"hello", "(up)"}

	got := projet.Erreur(input)
	want := []string{"HELLO"}

	for i, val := range got {
		if val != want[i] || len(got) > len(want) {
			t.Fatalf("Erreur()= %v, want %v", got, want)
		}
	}
}

func TestErreur_MultiUp(t *testing.T) {
	input := []string{"hello", "world","(up,","2)"}

	got := projet.Erreur(input)
	want := []string{"HELLO", "WORLD"}
	for i, val := range got {
		if val != want[i] || len(got) > len(want) {
			t.Fatalf("Erreur/Up = %v, want %v", got, want)
		}
	}
}

func TestPonctuation_MovesComma(t *testing.T) {
	input := []string{"hello", ",", "world"}

	got := projet.Ponctuation(input)
	want := []string{"hello,", "world"}

	for i, val := range got {
		if val != want[i] || len(got) > len(want) {
			t.Fatalf("Ponctuation()= %v, want %v", got, want)
		}
	}
}

func TestPonctuation_MissingComma(t *testing.T) {
	input := []string{"'hello", "world"}

	got := projet. RajouterApoGui(input,"'")
	want := []string{"'hello", "world'"}
	
	for i, val := range got {
		if val != want[i] || len(got) > len(want) {
			t.Fatalf("Ponctuation()= %v, want %v", got, want)
		}
	}
}

