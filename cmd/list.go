package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all contacts",
	Run: func(cmd *cobra.Command, args []string) {
		contacts, err := App.ListContacts()
		if err != nil {
			fmt.Println("Error listing contacts:", err)
			return
		}

		fmt.Println("\nContacts:")
		for _, c := range contacts {
			fmt.Printf("- [%d] %s <%s>\n", c.ID, c.Name, c.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
