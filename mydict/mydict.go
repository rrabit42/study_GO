package mydict

import "errors"

// map[string]string에 대한 alias(별명), struct 이 아님!
// type에 대해 alias를 명명할 수 있음!

// Dictionary type
type Dictionary map[string]string

// 이렇게 한꺼번에 선언 가능
var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-existing word")
	errCantDelete = errors.New("Cant delete non-existing word")
	errWordExists = errors.New("That word already exists")
)

// method를 type에도 추가할 수 있음

// Search for a word
func (d Dictionary) Search(word string) (string, error) {

	// map에서 key를 주면 값과 존재 여부(boolean)를 알려주는 좋은 코드을 제공함
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	// 단어가 있는지 확인 후 add
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	// 스위치가 더 나을듯
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }

	return nil
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	// 단어가 있는지 확인 후
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		// map delete는 이렇게!
		// delete function doesn't return anything, and will do notiong if the specified key doesn't exist
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
