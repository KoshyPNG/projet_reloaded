package projet

import "strings"

func Apostrophe(tab []string , word int , apostrophe bool,i int) ([]string ,bool)  {
	new := ""
	newtab := tab
	if i == 0 {
		if !apostrophe {
			if newtab[word] == "'" {
				new = strings.TrimPrefix(newtab[word], "'")
				newtab[word] = new
				if word < len(newtab) {
					newtab[word+1] += "'"
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
				newtab[word+1] += "'"
			}
			apostrophe = true
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
					newtab[word+1] += "\""
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
				newtab[word+1] += "\""
			}
			guillemet = true
		}
	}
	return newtab,guillemet
}