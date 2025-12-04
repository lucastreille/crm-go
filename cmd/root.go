package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/lucastreille/crm-go/internal/app"
	"github.com/lucastreille/crm-go/internal/notification"
	"github.com/lucastreille/crm-go/internal/storage"
)

var (
	cfgFile string
	App     *app.App
)

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "A simple CRM CLI",
	Long:  `A simple CRM CLI application to manage contacts using different storage backends.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.crm.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".crm")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	initApp()
}

func initApp() {
	var store storage.Storage
	var err error

	storageType := viper.GetString("storage.type")
	dbPath := viper.GetString("storage.path")

	switch storageType {
	case "json":
		if dbPath == "" {
			dbPath = "data/contacts.json"
		}
		store, err = storage.NewJSON(dbPath)
		if err != nil {
			fmt.Printf("Error initializing JSON storage: %v. Falling back to memory.\n", err)
			store = storage.NewMemory()
		} else {
			fmt.Println("Using JSON storage at", dbPath)
		}
	case "sqlite":
		if dbPath == "" {
			dbPath = "crm.db"
		}
		store, err = storage.NewGORMStore(dbPath)
		if err != nil {
			fmt.Printf("Error initializing SQLite storage: %v. Falling back to memory.\n", err)
			store = storage.NewMemory()
		} else {
			fmt.Println("Using SQLite storage at", dbPath)
		}
	default:
		fmt.Println("Using Memory storage (default)")
		store = storage.NewMemory()
	}

	notifiers := []notification.Notifier{
		notification.EmailNotifier{},
		notification.SmsNotifier{},
	}

	App = app.New(store, notifiers)
}
