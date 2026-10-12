package projet

import "strings"

func Ponctuation(tab []string) []string {
	return FormatPoint(tab)
}

func FormatPoint(tab []string) []string {
	apostrophe := false
	guillemet := false

	newtab := tab
	var new string
	
	for word := 0; word < len(newtab); word++ {
		if strings.HasPrefix(newtab[word], ".") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], ".")
				newtab[word] = new
				newtab[word-1] += "."
				word--
			}
		} else if strings.HasPrefix(newtab[word], ",") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], ",")
				newtab[word] = new
				newtab[word-1] += ","
				word--
			}
		} else if strings.HasPrefix(newtab[word], "!") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], "!")
				newtab[word] = new
				newtab[word-1] += "!"
				word--
			}
		} else if strings.HasPrefix(newtab[word], "?") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], "?")
				newtab[word] = new
				newtab[word-1] += "?"
				word--
			}
		} else if strings.HasPrefix(newtab[word], ":") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], ":")
				newtab[word] = new
				newtab[word-1] += ":"
				word--
			}
		} else if strings.HasPrefix(newtab[word], ";") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], ";")
				newtab[word] = new
				newtab[word-1] += ";"
				word--
			}
		} else if (strings.HasPrefix(newtab[word], "'")) || (strings.HasSuffix(newtab[word], "'"))  {
			if strings.HasPrefix(newtab[word], "'") {
				newtab,apostrophe = Apostrophe(newtab,word,apostrophe,0)
			} else {
				newtab,apostrophe = Apostrophe(newtab,word,apostrophe,1)
			}
		} else if (strings.HasPrefix(newtab[word], "\"")) || (strings.HasSuffix(newtab[word], "\""))  {
			if strings.HasPrefix(newtab[word], "\"") {
				newtab,guillemet = Guillemet(newtab,word,guillemet,0)
			} else {
				newtab,guillemet = Guillemet(newtab,word,guillemet,1)
			}
		}
	}
	newtab = Format(newtab)
	return newtab
}