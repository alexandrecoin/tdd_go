package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("Could not find the word you are looking for.")
	ErrWordExists = errors.New("This entry already exists")
)

func (d Dictionary) Search(entry string) (string, error) {
	definition, isWordFound := d[entry]
	if !isWordFound {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) AddEntry(entry, definition string) error {
	_, err := d.Search(entry)
	
	switch err {
	case ErrNotFound:
		d[entry] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	
	return nil
}
