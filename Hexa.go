package main

import "strconv"

func Hexa(sentence []string) []string {
	var result []string

	for _, h := range sentence {
		if h == "(hex)" {
			if len(result) > 0 { //ca change les groupe de caractères avant "(hex)"
				last := len(result) - 1

				n, err := strconv.ParseInt(result[last], 16, 64)// convertit le mot précédent de l'hexa vers un entier
				if err == nil { // si ce n'est pas de l'hexa valide, on laisse le mot tel quel
					result[last] = strconv.FormatInt(n, 10) // entier -> string en base 10
				}
			}
		}else {
			result = append(result, h) // mot normal : on le garde
		}
	}
	return result
}


