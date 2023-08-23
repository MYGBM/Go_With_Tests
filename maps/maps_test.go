package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want, "test")

	})
	t.Run("unknown word ", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertErrors(t, got, ErrNotFound)
	})

}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)
		assertErrors(t, err, nil) // we want error to be nil when adding a new word
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defintion := "this is just a test"
		dictionary := Dictionary{word: defintion}
		err := dictionary.Add(word, "new test")
		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, defintion)

	})
}
func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertErrors(t, err, ErrWordDoesNotExist)
	})
}
func assertStrings(t testing.TB, got, want, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q ", got, want, given)
	}
}
func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil { // word already exists in the dictionary
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition, word)
}
func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
