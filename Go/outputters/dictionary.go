package outputters

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrWordExists = DictionaryErr("ah ah ah, no duplicates allowed")
	ErrWordDoesNotExist = DictionaryErr("can't update something that doesn't exist")
	ErrNotFound = DictionaryErr("*waves hand* this isn't the word you're looking for")
)
func (d Dictionary) Search(word string) (string, error) {
	// Map lookups can return a boolean as a second value indicating if the key value was found
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)


	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	// Go has built-in delete function for maps
	delete(d, word)
}

func (e DictionaryErr) Error() string {
	return string(e)
}