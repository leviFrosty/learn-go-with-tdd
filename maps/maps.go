package maps

import "errors"

type Dictionary map[string]string

var (
	ErrorWordNotFound      = errors.New("could not find the word you were searching for")
	ErrorWordAlreadyExists = errors.New("could not add word, word already exists")
	ErrorWordDoesNotExist  = errors.New("cannot perform operation on word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorWordNotFound:
		d[word] = definition
	case nil:
		return ErrorWordAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorWordNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorWordNotFound:
		return ErrorWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
