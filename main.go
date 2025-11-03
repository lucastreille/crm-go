package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	Email string
}

var contacts = make(map[int]Contact)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) (string, error) {

	fmt.Print(prompt)

	t, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(t), nil

}

func readInt(prompt string) (int, error) {

	s, err := readLine(prompt)
	if err != nil {
		return 0, err
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("veuillez entrer un nombre valide")
	}

	return n, nil

}

func addContact(c Contact) error {

	if _, exists := contacts[c.ID]; exists {
		return fmt.Errorf("un contact avec l'ID %d existe déjà", c.ID)
	}
	contacts[c.ID] = c
	return nil

}

func listContacts() {

	if len(contacts) == 0 {
		fmt.Println("Aucun contact.")
		return
	}

	fmt.Println("Contacts:")
	for _, c := range contacts {
		fmt.Printf("- ID:%d | Nom:%s | Email:%s\n", c.ID, c.Name, c.Email)
	}

}

func updateContact(id int, name, email string) error {
	c, exists := contacts[id]
	if !exists {
		return fmt.Errorf("aucun contact avec l'ID %d", id)
	}
	if name != "" {
		c.Name = name
	}
	if email != "" {
		c.Email = email
	}
	contacts[id] = c
	return nil
}

func main() {

	for {

		fmt.Println("\n--- Mini-CRM ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact par ID")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")

		choice, err := readLine("Votre choix: ")
		if err != nil {
			fmt.Println("Erreur de lecture:", err)
			continue
		}

		switch choice {
		case "1":
			id, err := readInt("ID: ")

			if err != nil {
				fmt.Println("Erreur:", err)
				continue
			}

			name, err := readLine("Nom: ")
			if err != nil {
				fmt.Println("Erreur:", err)
				continue
			}

			email, err := readLine("Email: ")
			if err != nil {
				fmt.Println("Erreur:", err)
				continue
			}

			if err := addContact(Contact{ID: id, Name: name, Email: email}); err != nil {
				fmt.Println("Erreur:", err)
			} else {
				fmt.Println("Contact ajouté.")
			}

		case "2":
			listContacts()

		case "4":
			id, err := readInt("ID à mettre à jour: ")

			if err != nil {
				fmt.Println("Erreur:", err)
				continue
			}

			fmt.Println("(Laissez vide pour ne pas changer)")

			name, _ := readLine("Nouveau nom: ")
			email, _ := readLine("Nouvel email: ")
			if name == "" && email == "" {
				fmt.Println("Rien à mettre à jour.")
				continue
			}
			if err := updateContact(id, name, email); err != nil {
				fmt.Println("Erreur:", err)
			} else {
				fmt.Println("Contact mis à jour.")
			}

		case "5":
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide.")
		}

	}

}
