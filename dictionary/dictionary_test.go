package dictionary

import (
	"testing"
)

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func assertError(t testing.TB, got, want error)  {
	t.Helper()

	if got != want {
		t.Errorf("got error %q expected %q", got, want)
	}
	
}

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test":"this is a test"}

	t.Run("search for word in dict", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is a test"

		// if got != expected {
		// 	t.Errorf("got %q expected %q, given %q", got, expected, "test")
		// }
		assertStrings(t, got, expected)
	})

	t.Run("word not found in dict", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		// want := "word not found"

		// if err == nil {
		// 	t.Fatal("expected to get an error")
		// }

		// .Error() converts errors to a string, which is what we want when passing it to our assertion
		// assertStrings(t, err.Error(), want)
		
		assertError(t, err, ErrNotFound)
	})

}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string){
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should have found the word", err)
	}
	assertStrings(t, got, definition)
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is a test"

	dictionary.Add(word, definition)

	// got, err := dictionary.Search("test")
	// want := "this is a test"

	// if err != nil {
	// 	t.Fatal("should have found the added word,", err)
	// }
	// assertStrings(t, got, want)
	assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"

		dict := Dictionary{word: definition}

		err := dict.Add(word, "new definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})

}

func TestUpdate(t *testing.T){
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		newDef := "updated"

		dict := Dictionary{word: definition}

		err := dict.Update(word, newDef)
		
		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDef)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"

		dict := Dictionary{}

		err := dict.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T){
	word := "test"
	dict := Dictionary{word: "test definition"}

	dict.Delete(word)

	_, err := dict.Search(word)

	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}