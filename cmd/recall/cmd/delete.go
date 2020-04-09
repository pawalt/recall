package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(delete)
}

var delete = &cobra.Command{
	Use:   "delete",
	Short: "delete a command in the database",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("search requires exactly one argument, a name or substring of a command name. Received %d", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		deleteCmd(args[0])
	},
}

func deleteCmd(name string) {
	if i, err := strconv.Atoi(name); err == nil {
		err := store.Delete(i)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		entries, err := store.Search(name)
		if err != nil {
			log.Fatalln(err)
		}
		for _, entry := range entries {
			if strings.EqualFold(entry.Name, name) {
				err := store.Delete(entry.Index)
				if err != nil {
					log.Fatalln(err)
				}
				return
			}
		}
		fmt.Printf("Could not find a command name that matches \"%s\"\n", name)
	}
}
