package main

func Cut(data []byte) []string {
	var sentence []string
	var word []byte

	for _, s := range data {
		if s == ' ' {
			if len(word) > 0 { // évite les mots vides si plusieurs espaces
				sentence = append(sentence, string(word))
				word = nil
			}
		} else {
			word = append(word, s)
		}
	}

	// ajouter le dernier mot
	if len(word) > 0 {
		sentence = append(sentence, string(word))
	}

	return sentence
}