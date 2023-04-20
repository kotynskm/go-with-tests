package dictionary

type Dictionary map[string]string

// var (
// 	ErrNotFound = errors.New("word not found")
// 	ErrWordExists = errors.New("word already exists")
// ) 

// can refactor to make our errors const, and therefore more reusable and immutable
const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		// return "", errors.New("word not found")
		return "", ErrNotFound
	}
	return definition, nil
}

// func Search(dict map[string]string, key string) string {
// 	return dict[key]
// }

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error{
	
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

// since deleting a val that doesn't exist will have no effect, we don't need to complicate with error handling
func (d Dictionary) Delete(word string) {
	delete(d, word)
}