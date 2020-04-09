package datastore

// Entry is an individual name/command pairing
type Entry struct {
	Index   int
	Name    string
	Command string
}

// Datastore is an interface to allow you to store your entries
type Datastore interface {
	// Search returns all entries that contain name, case insensitive. Errors
	// out in case of i/o error
	Search(name string) ([]Entry, error)
	// Add adds a new entry with `name`. It returns an error in case of i/o
	// failure or if `name` already exists
	Add(name, command string) (int, error)
	// Update updates the entry at `index1. Error for i/o error or not exist
	Update(index int, command string) error
	// Get gets the entry at `index`, erroring out in case of i/o error. Returns nil
	// if `index` is not in the table
	Get(index int) (*Entry, error)
	// Delete deletes the entry at `index`, erroring out in case of i/o error or not found
	Delete(index int) error
}
