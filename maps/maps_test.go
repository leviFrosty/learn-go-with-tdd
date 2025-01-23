package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "This is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is just a test"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("x")
		if err == nil {
			t.Fatal("wanted error but didn't get one")
		}
		assertError(t, err, ErrorWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is a test"
		dict.Add(word, definition)
		assertDefinition(t, dict, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "word"
		definition := "this is a test"
		dict := Dictionary{word: definition}
		err := dict.Add("word", "something")
		assertError(t, err, ErrorWordAlreadyExists)
		assertDefinition(t, dict, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "word"
		dict := Dictionary{word: "foo"}
		definition := "bar"
		dict.Update(word, definition)
		assertDefinition(t, dict, word, definition)
	})
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update("word", "bar")
		assertError(t, err, ErrorWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "word"
		dict := Dictionary{word: "test definition"}
		err := dict.Delete(word)
		assertError(t, err, nil)
		_, err = dict.Search(word)
		assertError(t, err, ErrorWordNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "word"
		dict := Dictionary{}
		err := dict.Delete(word)
		assertError(t, err, ErrorWordDoesNotExist)
		_, err = dict.Search(word)
		assertError(t, err, ErrorWordNotFound)
	})
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	got, err := dict.Search(word)
	if err != nil {
		t.Fatalf("expected to find word")
	}
	assertString(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
