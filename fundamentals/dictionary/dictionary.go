package dictionary

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("cannot find word in dictionary")
	ErrWordExists       = DictionaryErr("word already exists in dictionary")
	ErrWordDoesNotExist = DictionaryErr("word does not exists in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	w, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return w, nil
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
	delete(d, word)
}
