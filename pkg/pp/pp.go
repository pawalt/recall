package pp

import (
	"fmt"
	"strconv"

	c "github.com/logrusorgru/aurora"
	"github.com/pawalt/recall/pkg/datastore"
)

// PrintEntry prints out a single datastore entry
func PrintEntry(entry *datastore.Entry) {
	fmt.Println(c.Green(entry.Name), "(#"+strconv.Itoa(entry.Index)+")")
	fmt.Println("  ", entry.Command)
}

// PrintEntries prints out a slice of datastore entries
func PrintEntries(entries []datastore.Entry) {
	for _, entry := range entries {
		PrintEntry(&entry)
	}
}
