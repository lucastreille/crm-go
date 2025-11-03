package main

import (
	"flag"
	"fmt"
)

func handleFlags() bool {
	addFlag := flag.Bool("add", false, "Ajouter un contact")
	idFlag := flag.Int("id", 0, "ID du contact")
	nameFlag := flag.String("name", "", "Nom du contact")
	emailFlag := flag.String("email", "", "Email du contact")
	flag.Parse()

	if *addFlag {
		if *idFlag == 0 || *nameFlag == "" || *emailFlag == "" {
			fmt.Println("Erreur: pour ajouter un contact, utilisez -id -name -email")
			return true
		}
		err := addContact(Contact{ID: *idFlag, Name: *nameFlag, Email: *emailFlag})
		if err != nil {
			fmt.Println("Erreur:", err)
		} else {
			fmt.Println("Contact ajouté via flags.")
		}
		return true
	}

	return false
}
