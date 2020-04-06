package datastore

// Entry is an individual name/command pairing
type Entry struct {
	Index   int
	Name    string
	Command string
}

// Datastore is an interface to allow you to store your entries
type Datastore interface {
	Search(name string) ([]*Entry, error)
	Add(name, command string) (int, error)
	Put(index int) error
	Get(index int) error
}
