package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(put)
}

var put = &cobra.Command{
	Use:   "put",
	Short: "put a command in the database",
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
		err := store.Put(i, command)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
