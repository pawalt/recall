package cmd

import (
	"fmt"
	"log"

	"github.com/pawalt/recall/pkg/pp"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(search)
}

var search = &cobra.Command{
	Use:   "search",
	Short: "search for a command in the database",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("search requires exactly one argument, a name or substring of a command name. Received %d", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		searchCmd(args[0])
	},
}

func searchCmd(name string) {
	entries, err := store.Search(name)
	if err != nil {
		log.Fatalln(err)
	}
	if len(entries) == 1 {
		fmt.Printf("1 result found!\n\n")
	} else {
		fmt.Printf("%d results found!\n\n", len(entries))
	}
	pp.PrintEntries(entries)
}
