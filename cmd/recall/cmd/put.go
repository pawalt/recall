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
	rootCmd.AddCommand(put)
}

var put = &cobra.Command{
	Use:   "put",
	Short: "add or update a command in the database",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("put requires exactly two arguments, a name/index and a command. Received %d", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		putCmd(args[0], args[1])
	},
}

func putCmd(name, command string) {
	if i, err := strconv.Atoi(name); err == nil {
		entry, err := store.Get(i)
		if err != nil {
			log.Fatalln(err)
		}
		if entry == nil {
			fmt.Printf("Could not find entry with index %d\n", i)
			os.Exit(1)
		} else {
			chooseOverwrite(entry, command)
		}
	} else {
		entries, err := store.Search(name)
		if err != nil {
			log.Fatalln(err)
		}
		if len(entries) == 0 {
			index, err := datastore.Add(store, name, command)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("Added at index %d\n", index)
		} else if len(entries) > 1 {
			fmt.Printf("Multiple matches for query \"%s\"\n", name)
			pp.PrintEntries(entries)
		} else {
			entry := &entries[0]
			if strings.EqualFold(entry.Name, name) {
				chooseOverwrite(&entries[0], command)
			} else {
				fmt.Printf("Only matched command has name \"%s\" but required name \"%s\".\n", entry.Name, name)
				fmt.Println("Requested name and name found in database must exactly match for deletion.")
				os.Exit(1)
			}
		}
	}
}

func chooseOverwrite(entry *datastore.Entry, command string) {
	choice := prompter.Choose(
		"overwrite \""+entry.Name+"\"?",
		[]string{"y", "n"},
		"n")
	if choice == "y" {
		entry.Command = command
		store.Put(*entry)
	} else {
		fmt.Println("Cancelled")
	}
}
