package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/codeskyblue/go-sh"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(run)
}

var run = &cobra.Command{
	Use:   "run",
	Short: "run a command in the database",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("run requires exactly one argument, a name or substring of a command name. Received %d", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		runCmd(args[0])
	},
}

func runCmd(name string) {
	if i, err := strconv.Atoi(name); err == nil {
		entry, err := store.Get(i)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Running command \"%s\"\n", entry.Name)
		sh.Command("sh", "-c", entry.Command).Run()
	} else {
		entries, err := store.Search(name)
		if err != nil {
			log.Fatalln(err)
		}
		for _, entry := range entries {
			if strings.EqualFold(entry.Name, name) {
				sh.Command("sh", "-c", entry.Command).Run()
				return
			}
		}
		fmt.Printf("Could not find a command name that matches \"%s\"\n", name)
	}
}
