// main.go
package main

import (
	"dict/packdict"
	"fmt"
)

func main() {
	dictionary := packdict.NewDictionary()

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
		nomToAdd := "ecole"
		definitionToAdd := "estiam"
		err = dictionary.Add(nomToAdd, definitionToAdd)
		if err != nil {
			fmt.Println("Error adding entry:", err)
		}
		fmt.Printf("Nom not found it's being added")
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
