package outputters

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("Error expected.")
		}

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "syzygy"
		definition := "straight-line configuration of three celestial bodies"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)

	})

	t.Run("existing word", func(t *testing.T) {
		word := "syzygy"
		definition := "straight-line configuration of three celestial bodies"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "mogrify"
		definition := "evil kung fu wizardry"
		dictionary := Dictionary{word: definition}
		newDefinition := "change shape incredibly"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "mogrify"
		definition := "evil kung fu wizardry"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "deletable"
	definition := "you know what this means"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}
// Helpers
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, definition)
}