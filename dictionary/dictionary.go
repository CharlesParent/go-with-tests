package dictionary

type Dictionary map[string]string

const (
	ErrorNotFound       = DictionaryErr("Could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("Word already exists")
	ErrWordDoesNotExist = DictionaryErr("Word does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		return ErrWordExists
	case ErrorNotFound:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = value
	case ErrorNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
