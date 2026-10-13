package main

import (
	"os"
	"projet"
	"strings"
	"fmt"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("il faut 2 fichier")
		return
	}

	entre := os.Args[1]
	sortie := os.Args[2]

	txt,err := os.ReadFile(entre)

	var tabbyte []byte
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	tabligne := strings.Split(string(txt), "\n")
	var oldphrase string
	var newphrase string
	for i, phrase := range tabligne {
		tabphrase := strings.Split(string(phrase), " ")
		tabphrase = projet.Format(tabphrase)

		for newphrase != oldphrase{
			oldphrase = strings.Join(tabphrase, " ")

			tabphrase = projet.Format(tabphrase)
			tabphrase = projet.Erreur(tabphrase)

			newphrase = strings.Join(tabphrase, " ")
		}
		
		tabphrase = projet.Erreur(tabphrase)
		tabligne[i] = ""
		for indice, mot := range tabphrase {
			tabligne[i] += mot
			if indice != len(tabphrase)-1 {
				tabligne[i] += " "
			}
		}
	}

	final := strings.Join(tabligne, "\n")
	tabbyte = []byte(final)
	err = os.WriteFile(sortie, tabbyte, os.FileMode(os.O_CREATE|os.O_WRONLY|os.O_TRUNC))
	if err != nil {
		fmt.Println("Erreur d'écriture", err)
		return
	}
}