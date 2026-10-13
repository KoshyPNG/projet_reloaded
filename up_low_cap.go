package projet

import (
	"fmt"
	"strconv"
	"strings"
)

func Convup(tab []string,rep string, i int) []string {
	nb,_ := strconv.Atoi(rep)
	if nb > len(tab[:i]){
		fmt.Println("up out of range")
		return Reecrire(tab, i)
	}
	for indice := 1; indice <= nb; indice++ {
		tab[i-indice] = strings.ToUpper(tab[i-indice])
	}
	return Reecrire(tab, i)
}

func Convlow(tab []string, rep string, i int) []string {
	nb,_ := strconv.Atoi(rep)
	if nb > len(tab[:i]){
		fmt.Println("low out of range")
		return Reecrire(tab, i)
	}
	for indice := 1; indice <= nb; indice++ {
		tab[i-indice] = strings.ToLower(tab[i-indice])
	}
	return Reecrire(tab, i)
}

func Convcap(tab []string, rep string, i int) []string {
	nb,_ := strconv.Atoi(rep)
	new := ""
	if nb > len(tab[:i]){
		fmt.Println("up out of range")
		return Reecrire(tab, i)
	}
	for indice := 1; indice <= nb; indice++ {
		for i, lettre := range tab[i-indice] {
			if i == 0 {
				new += strings.ToUpper(string(lettre))
			} else {
				new += string(lettre)
			}
		}
		tab[i-indice] = new
		new = ""
	}
	return Reecrire(tab, i)
}