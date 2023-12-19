package main

import (
	"dict/packdict"
	"fmt"
)

func main() {
	dictionary := packdict.NewDictionary()

	dictionary.Add("estiam", "ecole")
	dictionary.Add("go", "language")

	fmt.Println("avant remove:")
	dictionary.List()

	keyToFind := "go"
	if value, found := dictionary.Get(keyToFind); found {
		fmt.Printf("Value found for key %s: %s\n", keyToFind, value)
	} else {
		fmt.Printf("Key not found: %s\n", keyToFind)
	}

	keyToRemove := "go"
	dictionary.Remove(keyToRemove)

	fmt.Println("apres remove:")
	dictionary.List()
}
