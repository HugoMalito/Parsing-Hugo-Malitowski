package main

import "strconv"

func NonBinaire(sentence []string) []string {
	var result []string

	for _, b := range sentence {
		if b == "(bin)" {
			if len(result) > 0 { //ca change les groupe de caractères avant "(bin)"
				last := len(result) - 1

				n, err := strconv.ParseInt(result[last], 2, 64)// convertit le mot précédent du binaire vers un entier
				if err == nil { // si ce n'est pas de l'hexa valide, on laisse le mot tel quel
					result[last] = strconv.FormatInt(n, 10) // entier -> string en base 10
				}
			}
		}else {
			result = append(result, b) // mot normal : on le garde
		}
	}
	return result
}


