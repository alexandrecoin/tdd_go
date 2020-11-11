package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, isWordFound := d[word]
	if !isWordFound {
		return "", errors.New("Could not find the word you are looking for.")
	}
	return definition, nil
}