package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("Could not find the word you are looking for.")

func (d Dictionary) Search(word string) (string, error) {
	definition, isWordFound := d[word]
	if !isWordFound {
		return "", ErrNotFound
	}
	return definition, nil
}