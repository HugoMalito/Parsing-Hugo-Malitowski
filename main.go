package main

import (
	"log"
	"os"
	"strings"
)

const (
	carspe = ".,?;/:§!µ*%£$_-&=+°`|"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	mots := Cut(data)
	mots = Hexa(mots)
	mots = LeMaxeur(mots)
	mots = MartiMinimus(mots)
	mots = Capitalisme(mots)
	mots = NonBinaire(mots)
	mots = Anal_yse(mots)
	texte := strings.Join(mots, " ")
	texte = PointSurLesI(texte)

	err = os.WriteFile("output.txt", []byte(texte), 0644)
	if err != nil {
		log.Fatal(err)
	}
}