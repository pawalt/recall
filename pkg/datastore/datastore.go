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
	// Put inserts an entry into the datastore. error on i/o error
	Put(entry Entry) error
	// Get gets the entry at `index`, erroring out in case of i/o error. Returns nil
	// if `index` is not in the table
	Get(index int) (*Entry, error)
	// Delete deletes the entry at `index`, erroring out in case of i/o error or not found
	Delete(index int) error
	// FreshIndex returns an unused index in the datastore, error on i/o error
	FreshIndex() (int, error)
}

// Add adds a new value to the datastore ds
func Add(ds Datastore, name, command string) (int, error) {
	index, err := ds.FreshIndex()
	if err != nil {
		return -1, err
	}
	entry := Entry{
		Index:   index,
		Name:    name,
		Command: command,
	}
	err = ds.Put(entry)
	if err != nil {
		return -1, err
	}
	return index, nil
}
