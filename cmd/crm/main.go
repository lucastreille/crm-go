package main

import (
	"fmt"

	"github.com/lucastreille/crm-go/internal/app"
	"github.com/lucastreille/crm-go/internal/storage"
)

func main() {
	store := storage.NewMemory()
	application := app.New(store)

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("0. Quitter")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			var id int
			var name, email string
			fmt.Print("ID : ")
			fmt.Scan(&id)
			fmt.Print("Nom : ")
			fmt.Scan(&name)
			fmt.Print("Email : ")
			fmt.Scan(&email)

			if err := application.AddContact(id, name, email); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Contact ajouté !")
			}

		case 2:
			contacts, _ := application.ListContacts()
			if len(contacts) == 0 {
				fmt.Println("Aucun contact.")
				continue
			}
			for _, c := range contacts {
				fmt.Printf("ID: %d | Nom: %s | Email: %s\n", c.ID, c.Name, c.Email)
			}

		case 3:
			var id int
			fmt.Print("ID du contact à supprimer : ")
			fmt.Scan(&id)
			if err := application.DeleteContact(id); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Contact supprimé.")
			}

		case 4:
			var id int
			var newName, newEmail string
			fmt.Print("ID du contact : ")
			fmt.Scan(&id)
			fmt.Print("Nouveau nom : ")
			fmt.Scan(&newName)
			fmt.Print("Nouvel email : ")
			fmt.Scan(&newEmail)

			if err := application.UpdateContact(id, newName, newEmail); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Contact mis à jour.")
			}

		case 0:
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
