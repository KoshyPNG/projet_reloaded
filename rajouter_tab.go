package projet

func Rajouter(tab []string , i int , mot string) []string {
	var slice []string
	for indice, val := range tab {
		if indice == i+1 {
			slice = append(slice, mot)
		}
		slice = append(slice, val)
	}
	return slice 
}