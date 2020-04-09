package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/pawalt/recall/pkg/datastore"
	"github.com/pawalt/recall/pkg/pp"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rename)
}

var rename = &cobra.Command{
	Use:   "rename",
	Short: "rename a command in the database",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("rename requires exactly two arguments, a name/index and a new name. Received %d", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		renameCmd(args[0], args[1])
	},
}

func renameCmd(oldName, newName string) {
	if i, err := strconv.Atoi(oldName); err == nil {
		entry, err := store.Get(i)
		if err != nil {
			log.Fatalln(err)
		}
		chooseRename(entry, newName)
	} else {
		entries, err := store.Search(oldName)
		if err != nil {
			log.Fatalln(err)
		}
		if len(entries) == 0 {
			fmt.Printf("Could not find match for query %s\n", oldName)
			os.Exit(1)
		} else if len(entries) > 1 {
			fmt.Printf("Multiple matches for query \"%s\"\n", oldName)
			pp.PrintEntries(entries)
		} else {
			entry := &entries[0]
			if strings.EqualFold(entry.Name, oldName) {
				chooseRename(&entries[0], newName)
			} else {
				fmt.Printf("Only matched command has name \"%s\" but required name \"%s\".\n", entry.Name, oldName)
				fmt.Println("Requested name and name found in database must exactly match for deletion.")
				os.Exit(1)
			}
		}
	}
}

func chooseRename(entry *datastore.Entry, newName string) {
	choice := prompter.Choose(
		"rename \""+entry.Name+"\"?",
		[]string{"y", "n"},
		"n")
	if choice == "y" {
		entry.Name = newName
		store.Put(*entry)
	} else {
		fmt.Println("Cancelled")
	}
}
