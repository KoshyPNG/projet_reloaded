package projet

import (
	"strings"
)

func RajouterApoGui(tab []string, lettre string) []string {
	new := strings.TrimPrefix(tab[len(tab)-1], " ") + lettre
	tab[len(tab)-1] = new
	return tab
}

func Apostrophe(tab []string , word int , apostrophe bool,i int) ([]string ,bool)  {
	new := ""
	newtab := tab
	if i == 0 {
		if !apostrophe {
			if newtab[word] == "'" {
				new = strings.TrimPrefix(newtab[word], "'")
				newtab[word] = new
				if word < len(newtab)-1 {
					new = "'" + strings.TrimPrefix(newtab[word+1], " ")
					newtab[word+1] = new
				}
			}
			apostrophe = true
		} else {
			new = strings.TrimPrefix(newtab[word], "'")
			newtab[word] = new
			newtab[word-1] += "'"
			apostrophe = false
		}
	} else {
		if !apostrophe {
			if newtab[word] == "'" {
				new = strings.TrimSuffix(newtab[word], "'")
				newtab[word] = new
				newtab[word-1] += "'"
			}
			apostrophe = true
		} else {
			if newtab[word] == "'" {
				new = strings.TrimPrefix(newtab[word], "'")
				newtab[word] = new
				newtab[word-1] += "'"
			}
			apostrophe = false
			
		}
	}
	return newtab,apostrophe
}

func Guillemet(tab []string , word int , guillemet bool,i int) ([]string ,bool)  {
	new := ""
	newtab := tab
	if i == 0 {
		if !guillemet {
			if newtab[word] == "\"" {
				new = strings.TrimPrefix(newtab[word], "\"")
				newtab[word] = new
				if word < len(newtab) {
					new = "\"" +  strings.TrimPrefix(newtab[word+1], " ")
					newtab[word+1] = new
				}
			}
			guillemet = true
		} else {
			new = strings.TrimPrefix(newtab[word], "\"")
			newtab[word] = new
			newtab[word-1] += "\""
			guillemet = false
		}
	} else {
		if !guillemet {
			if newtab[word] == "\"" {
				new = strings.TrimSuffix(newtab[word], "\"")
				newtab[word] = new
				newtab[word-1] += "\""
			}
			guillemet = true
		} else {
			if newtab[word] == "\"" {
				new = strings.TrimPrefix(newtab[word], "\"")
				newtab[word] = new
				newtab[word-1] += "\""
			}
			guillemet = true
		}
	}
	return newtab,guillemet
}