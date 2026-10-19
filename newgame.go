package main

import (
	"fmt"
)

func main() {
	newgames()
}

func newgames() {
	fmt.Println("Bienvenue dans le jeu du pendu ! en go")
	var pseudo string
	fmt.Print("Entrez votre pseudo : ")
	fmt.Scanln(&pseudo)

	if pseudo == "" {
		fmt.Println("Erreur : Entrez un pseudo valide")
		return
	}
	fmt.Println("Bienvenue", pseudo)

	mot := selection()
	if mot == "" {
		return
	}
	fmt.Println("Nouvelle partie ! le mot contient", len(mot), "lettres")

	IndicateLetter(mot)
}
