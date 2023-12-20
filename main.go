// main.go
package main

import (
	"dict/packdict"
	"fmt"
)

func main() {
	dictionary := packdict.NewDictionary()

	err := dictionary.Add("exemple", "exemple def")
	err = dictionary.Add("platform", "teams")
	err = dictionary.Add("formateur", "Aziz")

	fmt.Println("Before remove:")
	dictionary.List()

	nomToFind := "estiam"
	if definition, found := dictionary.Get(nomToFind); found {
		fmt.Printf("Definition found for nom %s: %s\n", nomToFind, definition)
	} else {
		nomToAdd := "ecole"
		definitionToAdd := "estiam"
		err = dictionary.Add(nomToAdd, definitionToAdd)
		fmt.Printf("Nom not found but it's being added")
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

	nomToUpdate := "formateur"
	newDefinition := "updated definition"
	err = dictionary.Update(nomToUpdate, newDefinition)
	if err != nil {
		fmt.Println("Error updating entry:", err)
	}

	fmt.Println("After update:")
	dictionary.List()
}
