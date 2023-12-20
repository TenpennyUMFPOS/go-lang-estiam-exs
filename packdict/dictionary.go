// dictionary.go
package packdict

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type DictionaryEntry struct {
	Nom        string `json:"nom"`
	Definition string `json:"definition"`
}

type Dictionary struct {
	Entries []DictionaryEntry
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		Entries: make([]DictionaryEntry, 0),
	}
}

func (d *Dictionary) Get(nom string) (string, bool) {
	for _, entry := range d.Entries {
		if entry.Nom == nom {
			return entry.Definition, true
		}
	}
	return "", false
}

func (d *Dictionary) Add(nom, definition string) error {
	entry := DictionaryEntry{Nom: nom, Definition: definition}
	d.Entries = append(d.Entries, entry)

	err := d.saveToJSON()
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionary) Remove(nom string) error {
	var index int
	for i, entry := range d.Entries {
		if entry.Nom == nom {
			index = i
			break
		}
	}
	d.Entries = append(d.Entries[:index], d.Entries[index+1:]...)

	err := d.saveToJSON()
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionary) List() {
	sort.Slice(d.Entries, func(i, j int) bool {
		return d.Entries[i].Nom < d.Entries[j].Nom
	})

	for _, entry := range d.Entries {
		fmt.Printf("%s: %s\n", entry.Nom, entry.Definition)
	}
}

func (d *Dictionary) LoadFromJSON(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &d.Entries)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionary) saveToJSON() error {
	data, err := json.MarshalIndent(d.Entries, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("details.json", data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionary) Update(nom, newDefinition string) error {
	for i, entry := range d.Entries {
		if entry.Nom == nom {
			d.Entries[i].Definition = newDefinition

			err := d.saveToJSON()
			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("Entry not found with name: %s", nom)
}
