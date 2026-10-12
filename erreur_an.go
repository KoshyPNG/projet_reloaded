package projet

import "strings"

func ErrAn(tab []string , i int) []string {
	newtab := tab
	if i < len(newtab)-1 {
		switch newtab[i]{
		case "A":
			if Aslettre(newtab,i+1) {
				newtab[i] = "An"
			}
		case "a":
			if Aslettre(newtab,i+1) {
				newtab[i] = "an"
			}
		case "An":
			if !Aslettre(newtab,i+1) {
				newtab[i] = "A"
			}
		case "an":
			if !Aslettre(newtab,i+1) {
				newtab[i] = "an"
			}
		}
	}
	return newtab
}

func Aslettre(tab []string , i int) bool {
	first := strings.ToLower(string(tab[i][0]))
	return strings.Contains("aeiouh", first)
}