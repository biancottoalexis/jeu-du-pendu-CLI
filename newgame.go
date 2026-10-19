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
	fmt.Println("Bienvenue", pseudo)

	if pseudo == "" {
		fmt.Println("Erreur : Entrez un pseudo valide")
		return
	}
	mot := selection()
	fmt.Println("Nouvelle partie ! le mot contient", len(mot), "lettres")
}
