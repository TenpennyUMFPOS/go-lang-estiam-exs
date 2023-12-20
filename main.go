// main.go
package main

import (
	"dict/packdict"
	"fmt"
	"sync"
)

func main() {
	dictionary := packdict.NewDictionary()

	//Start chanel
	addChannel := make(chan packdict.DictionaryEntry)

	var wg sync.WaitGroup
	go func() {
		for entry := range addChannel {
			dictionary.Add(entry.Nom, entry.Definition)

		}
	}()

	entriesToAdd := []packdict.DictionaryEntry{
		{Nom: "exemple", Definition: "exemple def"},
		{Nom: "platform", Definition: "teams"},
		{Nom: "formateur", Definition: "Aziz"},
	}

	for _, entry := range entriesToAdd {
		wg.Add(1)
		go func(e packdict.DictionaryEntry) {
			defer wg.Done()
			fmt.Println("chanel working ..")
			fmt.Printf("Sending entry: %s\n", e.Nom)
			// Send the entry to the channel
			addChannel <- e
		}(entry)
	}

	wg.Wait()

	close(addChannel)

	// channel

	err := dictionary.Add("go", "language de programation")

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
