package main

type Dictionary map[string]string

const (
	ErrNotFound = ErrDictionary("Could not find the word you are looking for.")
	ErrWordExists = ErrDictionary("This entry already exists")
	ErrWordDoesNotExist = ErrDictionary("This entry does not exist. Can't update.")
)

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

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

func (d Dictionary) UpdateEntry(entry, newDefinition string) error{
	_, err := d.Search(entry)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[entry] = newDefinition
	default:
		return err	
	}
	
	return nil
}
