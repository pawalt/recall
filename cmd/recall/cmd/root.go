package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pawalt/recall/pkg/datastore"
	"github.com/spf13/cobra"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)
var store datastore.Datastore

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	storePath := filepath.Join(home, ".config", "recall", "data.json")
	store, err = datastore.NewJSONDatastore(storePath)
	if err != nil {
		log.Fatalln(err)
	}
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
