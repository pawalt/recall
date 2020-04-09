package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pawalt/recall/pkg/datastore"
	"github.com/spf13/cobra"
)

var (
	store    datastore.Datastore
	dataFile string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&dataFile, "data", "", "data file (default is $HOME/.config/recall/data.json)")
}

var rootCmd = &cobra.Command{
	Use:   "recall",
	Short: "recall is a tool to help you remember that pesky command",
	Long:  `For more information, visit https://github.com/pawalt/recall`,
}

// Execute starts the cobra chain
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if dataFile == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
		dataFile = filepath.Join(home, ".config", "recall", "data.json")
	}
	var err error
	store, err = datastore.NewJSONDatastore(dataFile)
	if err != nil {
		log.Fatalln(err)
	}
}
