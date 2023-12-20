// main.go
package main

import (
	"dict/packdict"
	"fmt"
)

func main() {
	dictionary := packdict.NewDictionary()

	// Add some entries
	err := dictionary.Add("estiam", "ecole")
	if err != nil {
		fmt.Println("Error adding entry:", err)
	}

	err = dictionary.Add("platform", "teams")
	if err != nil {
		fmt.Println("Error adding entry:", err)
	}

	err = dictionary.Add("formateur", "Aziz")
	if err != nil {
		fmt.Println("Error adding entry:", err)
	}

	fmt.Println("Before remove:")
	dictionary.List()

	nomToFind := "estiam"
	if definition, found := dictionary.Get(nomToFind); found {
		fmt.Printf("Definition found for nom %s: %s\n", nomToFind, definition)
	} else {
		fmt.Printf("Nom not found: %s\n", nomToFind)
	}

	nomToAdd := "new_nom"
	definitionToAdd := "new_definition"
	err = dictionary.Add(nomToAdd, definitionToAdd)
	if err != nil {
		fmt.Println("Error adding entry:", err)
	}

	fmt.Println("After add:")
	dictionary.List()

	nomToRemove := "estiam"
	err = dictionary.Remove(nomToRemove)
	if err != nil {
		fmt.Println("Error removing entry:", err)
	}

	fmt.Println("After remove:")
	dictionary.List()
}
