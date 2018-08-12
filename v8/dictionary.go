package dictionary

//Dictionary type
type Dictionary map[string]string

//DictionaryErr error
type DictionaryErr string

//ErrNotFound error
const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot update or delete word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

//Search test
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

//Add add a word to dictionary
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

//Update word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

//Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
