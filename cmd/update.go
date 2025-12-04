package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	updateID    int
	updateName  string
	updateEmail string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing contact",
	Run: func(cmd *cobra.Command, args []string) {
		if updateID == 0 {
			fmt.Println("Error: ID is required for update")
			return
		}

		err := App.UpdateContact(updateID, updateName, updateEmail)
		if err != nil {
			fmt.Println("Error updating contact:", err)
		} else {
			fmt.Println("Contact updated successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVar(&updateID, "id", 0, "Contact ID (required)")
	updateCmd.Flags().StringVar(&updateName, "name", "", "New Name")
	updateCmd.Flags().StringVar(&updateEmail, "email", "", "New Email")
	updateCmd.MarkFlagRequired("id")
}
