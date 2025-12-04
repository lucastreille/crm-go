package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	addID    int
	addName  string
	addEmail string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new contact",
	Run: func(cmd *cobra.Command, args []string) {
		if addName == "" || addEmail == "" {
			// Interactive mode
			fmt.Println("Interactive mode:")
			if addID == 0 {
				fmt.Print("ID: ")
				fmt.Scan(&addID)
			}
			if addName == "" {
				fmt.Print("Name: ")
				fmt.Scan(&addName)
			}
			if addEmail == "" {
				fmt.Print("Email: ")
				fmt.Scan(&addEmail)
			}
		}

		err := App.AddContact(addID, addName, addEmail)
		if err != nil {
			fmt.Println("Error adding contact:", err)
		} else {
			fmt.Println("Contact added successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVar(&addID, "id", 0, "Contact ID")
	addCmd.Flags().StringVar(&addName, "name", "", "Contact Name")
	addCmd.Flags().StringVar(&addEmail, "email", "", "Contact Email")
}
