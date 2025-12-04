package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a contact",
	Run: func(cmd *cobra.Command, args []string) {
		if deleteID == 0 {
			fmt.Println("Error: ID is required for deletion")
			return
		}

		err := App.DeleteContact(deleteID)
		if err != nil {
			fmt.Println("Error deleting contact:", err)
		} else {
			fmt.Println("Contact deleted successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVar(&deleteID, "id", 0, "Contact ID (required)")
	deleteCmd.MarkFlagRequired("id")
}
