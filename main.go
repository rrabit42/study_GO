package main

import (
	"fmt"

	"github.com/rrabit42/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "Firstword"}

	word := "hello"
	definition := "Greeting"

	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}

	def, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("found", word, "definition:", def)
	}

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := dictionary.Update(word, "Second")
	if err3 != nil {
		fmt.Println(err3)
	}
	baseWord, _ := dictionary.Search(word)
	fmt.Println(baseWord)

	err4 := dictionary.Delete(word)
	if err4 != nil {
		fmt.Println(err4)
	}

	def, err5 := dictionary.Search(word)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println("found", word, "definition:", def)
	}
}
