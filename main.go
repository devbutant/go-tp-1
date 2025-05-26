package main

import (
	"flag"
	"fmt"

	"github.com/devbutant/tp-note/annuaire"
)

func main() {
	// Initialisation des contacts par défaut
	annuaire.Init()

	action := flag.String("action", "", "Action à effectuer: ajouter, rechercher, lister, supprimer, modifier")
	nom := flag.String("nom", "", "Nom du contact")
	tel := flag.String("tel", "", "Numéro de téléphone")

	flag.Parse()

	switch *action {
	case "ajouter":
		if annuaire.Ajouter(*nom, *tel) {
			fmt.Println("Contact ajouté avec succès.")
		} else {
			fmt.Println("Contact déjà existant.")
		}
	case "rechercher":
		if c, ok := annuaire.Rechercher(*nom); ok {
			fmt.Printf("Nom: %s, Téléphone: %s\n", c.Nom, c.Tel)
		} else {
			fmt.Println("Contact non trouvé.")
		}
	case "lister":
		for _, c := range annuaire.Lister() {
			fmt.Printf("Nom: %s, Téléphone: %s\n", c.Nom, c.Tel)
		}
	case "modifier":
		if annuaire.Modifier(*nom, *tel) {
			fmt.Println("Contact modifié.")
		} else {
			fmt.Println("Contact non trouvé.")
		}
	case "supprimer":
		if annuaire.Supprimer(*nom) {
			fmt.Println("Contact supprimé.")
		} else {
			fmt.Println("Contact non trouvé.")
		}
	default:
		fmt.Println("Action non reconnue.")
	}
}
