// main.go
package main

import (
	"dict/packdict"
	"fmt"
	"net/http"
	"sort"
)

func main() {
	dictionary := packdict.NewDictionary()

	err := dictionary.LoadFromJSON("details.json")
	if err != nil {
		fmt.Println("Error loading data from JSON:", err)
	}

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		nom := r.URL.Query().Get("nom")
		definition, found := dictionary.Get(nom)
		if found {
			fmt.Fprintf(w, "Definition for %s: %s", nom, definition)
		} else {
			http.Error(w, "Entry not found", http.StatusNotFound)
		}
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		nom := r.URL.Query().Get("nom")
		definition := r.URL.Query().Get("definition")
		err := dictionary.Add(nom, definition)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error adding entry: %s", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Entry added successfully")
	})

	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		nom := r.URL.Query().Get("nom")
		err := dictionary.Remove(nom)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error removing entry: %s", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Entry removed successfully")
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		sort.Slice(dictionary.Entries, func(i, j int) bool {
			return dictionary.Entries[i].Nom < dictionary.Entries[j].Definition
		})
		for _, entry := range dictionary.Entries {
			fmt.Fprintf(w, "%s: %s \n", entry.Nom, entry.Definition)
		}
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		nom := r.URL.Query().Get("nom")
		newDefinition := r.URL.Query().Get("newDefinition")
		err := dictionary.Update(nom, newDefinition)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error updating entry: %s", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Entry updated successfully")
	})

	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
