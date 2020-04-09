package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// JSONDatastore is a simple JSON list representation of a datastore
type JSONDatastore struct {
	store *JSONStorage
	path  string
}

// JSONStorage is the storage format for the JSON file
type JSONStorage struct {
	Entries []Entry
}

// NewJSONDatastore loads in the JSON list at `path`
func NewJSONDatastore(path string) (*JSONDatastore, error) {
	// If path does not exist, create it and populate it with blank entry
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			dir := filepath.Dir(path)
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				log.Fatalln(err)
			}
			store := &JSONStorage{
				Entries: []Entry{},
			}
			bytes, err := json.Marshal(store)
			if err != nil {
				return nil, err
			}
			err = ioutil.WriteFile(path, bytes, 0644)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	store := JSONStorage{}
	err = json.Unmarshal(bytes, &store)
	if err != nil {
		return nil, err
	}

	return &JSONDatastore{
		store: &store,
		path:  path,
	}, nil
}

// Search fulfills the datastore interface
func (j *JSONDatastore) Search(name string) ([]Entry, error) {
	res := make([]Entry, 0, 10)
	for _, item := range j.store.Entries {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(name)) {
			res = append(res, item)
		}
	}
	return res, nil
}

// Add fulfills the datastore interface
func (j *JSONDatastore) Add(name, command string) (int, error) {
	max := 1
	for _, item := range j.store.Entries {
		if item.Name == name {
			return -1, errors.New("already have item with name " + item.Name)
		}
		if item.Index > max {
			max = item.Index
		}
	}

	entry := Entry{
		Index:   max,
		Name:    name,
		Command: command,
	}

	err := j.addEntry(entry)
	if err != nil {
		return -1, err
	}

	return max, nil
}

// Update fulfills the datastore interface
func (j *JSONDatastore) Update(index int, command string) error {
	for _, item := range j.store.Entries {
		if item.Index == index {
			item.Command = command
			return nil
		}
	}

	return fmt.Errorf("could not find item at index %d", index)
}

// Get fulfills the datastore interface
func (j *JSONDatastore) Get(index int) (*Entry, error) {
	for _, item := range j.store.Entries {
		if item.Index == index {
			copy := item
			return &copy, nil
		}
	}

	return nil, nil
}

// Delete fulfills the datastore interface
func (j *JSONDatastore) Delete(index int) error {
	for sliceInd, item := range j.store.Entries {
		if item.Index == index {
			return j.removeEntry(sliceInd)
		}
	}

	return fmt.Errorf("could not find item at index %d", index)
}

func (j *JSONDatastore) removeEntry(index int) error {
	newStore := JSONStorage{
		Entries: append(j.store.Entries[:index], j.store.Entries[index+1:]...),
	}

	bytes, err := json.Marshal(newStore)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(j.path, bytes, 0644)
	if err != nil {
		return err
	}

	j.store = &newStore

	return nil
}

func (j *JSONDatastore) addEntry(entry Entry) error {
	newStore := JSONStorage{
		Entries: append(j.store.Entries, entry),
	}

	bytes, err := json.Marshal(newStore)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(j.path, bytes, 0644)
	if err != nil {
		return err
	}

	j.store = &newStore

	return nil
}
