package projet

import "strings"

func Erreur(tab []string) []string {
	newtab := tab
	new := ""
	for word := 0; word < len(newtab); word++ {
		if (strings.HasPrefix(newtab[word],"(hex)")) || (strings.HasSuffix(newtab[word],"(hex)")) {
			if strings.HasPrefix(newtab[word],"(hex)") {
				new = strings.TrimPrefix(newtab[word],"(hex)")
			} else {
				new = strings.TrimSuffix(newtab[word],"(hex)")
			}
			newtab[word] = "(hex)"
			if new != "" {
				newtab = Rajouter(newtab,word,new)
			}
			newtab = Convhexa(newtab, word)
			word--

		} else if (strings.HasPrefix(newtab[word],"(bin)")) || (strings.HasSuffix(newtab[word],"(bin)")) {
			if strings.HasPrefix(newtab[word],"(bin)") {
				new = strings.TrimPrefix(newtab[word],"(bin)")
			} else {
				new = strings.TrimSuffix(newtab[word],"(bin)")
			}
			newtab[word] = "(bin)"
			if new != "" {
				newtab = Rajouter(newtab,word,new)
			}
			newtab = Convbin(newtab, word)
			word--

		} else if (strings.HasPrefix(newtab[word],"(up)")) || (strings.HasSuffix(newtab[word],"(up)")) {
			if strings.HasPrefix(newtab[word],"(up)") {
				new = strings.TrimPrefix(newtab[word],"(up)")
			} else {
				new = strings.TrimSuffix(newtab[word],"(up)")
			}
			newtab[word] = "(up)"
			if new != "" {
				newtab = Rajouter(newtab,word,new)
			}
			newtab = Convup(newtab, "1", word)
			word--

		} else if (strings.HasSuffix(newtab[word],"(up,")) {
			new = strings.TrimSuffix(newtab[word],"(up)")
			newtab[word] = "(up,"
			rep := Repetion(newtab[word+1])
			newtab = Reecrire(newtab, word+1)
			if new != "" {
				newtab = Rajouter(newtab,word,new)
			}
			newtab = Convup(newtab, rep, word)
			word--
			word--
		} else if (strings.HasPrefix(newtab[word],"(low)")) || (strings.HasSuffix(newtab[word],"(low)")) {
			if strings.HasPrefix(newtab[word],"(low)") {
				new = strings.TrimPrefix(newtab[word],"(low)")
			} else {
				new = strings.TrimSuffix(newtab[word],"(low)")
			}
			newtab[word] = "(low)"
			if new != "" {
				newtab = Rajouter(newtab,word,new)
			}
			newtab = Convlow(newtab, "1", word)
			word--

		} else if (strings.HasSuffix(newtab[word],"(low,")) {
			new = strings.TrimSuffix(newtab[word],"(low,")
			newtab[word] = "(low,"
			rep := Repetion(newtab[word+1])
			newtab = Reecrire(newtab, word+1)
			if new != "" {
				newtab = Rajouter(newtab,word,new)
			}
			newtab = Convlow(newtab, rep, word)
			word--
			word--
		} else if (strings.HasPrefix(newtab[word],"(cap)")) || (strings.HasSuffix(newtab[word],"(cap)")) {
			if strings.HasPrefix(newtab[word],"(cap)") {
				new = strings.TrimPrefix(newtab[word],"(cap)")
			} else {
				new = strings.TrimSuffix(newtab[word],"(cap)")
			}
			newtab[word] = "(cap)"
			if new != "" {
				newtab = Rajouter(newtab,word,new)
			}
			newtab = Convcap(newtab, "1", word)
			word--

		} else if (strings.HasSuffix(newtab[word],"(cap,")) {
			new = strings.TrimSuffix(newtab[word],"(cap)")
			newtab[word] = "(cap,"
			rep := Repetion(newtab[word+1])
			newtab = Reecrire(newtab, word+1)
			if new != "" {
				newtab = Rajouter(newtab,word,new)
			}
			newtab = Convcap(newtab, rep, word)
			word--
			word--
		}
		newtab = ErrAn(newtab,word)
	}
	
	newtab = Format(newtab)
	newtab = Ponctuation(newtab)
	
	return newtab
}

func Repetion(word string) string {
	for _, lettre := range word {
		if lettre >= '1' && lettre <= '9' {
			return string(lettre)
		}
	}
	return "1"
}