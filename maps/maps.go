package main

const (
	ErrNotFound   = DictionaryErr("couldn't find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string // creating our own error type

func (e DictionaryErr) Error() string { // our error type implements the error interface
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound: //word doesn't exist
		d[word] = definition
	case nil: //NO error from search, the word could be found
		return ErrWordExists
	default:
		return err
	}
	return nil
}
