package main

func isPunct(c byte) bool {
	return c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';'
}

func PointSurLesI(s string) string {
	var result []byte
	i := 0

	for i < len(s) {
		if !isPunct(s[i]) {
			result = append(result, s[i])
			i++
			continue
		}
		j := i
		for j < len(s) && isPunct(s[j]) {
			j++
		}
		for len(result) > 0 && result[len(result)-1] == ' ' { //retire les espaces avant
			result = result[:len(result)-1]
		}
		result = append(result, s[i:j]...) // met un espace après s'il y a un mot
		if j < len(s) && s[j] != ' ' {
			result = append(result, ' ')
		}
		i = j
	}

	return string(result)
}
