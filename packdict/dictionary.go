package packdict

import (
	"fmt"
	"sort"
)

type Dictionary struct {
	m map[string]string
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		m: make(map[string]string),
	}
}

func (d *Dictionary) Get(key string) (string, bool) {
	value, found := d.m[key]
	return value, found
}

func (d *Dictionary) Add(key string, value string) {
	d.m[key] = value
}

func (d *Dictionary) Remove(key string) {
	delete(d.m, key)
}

func (d *Dictionary) List() {
	var keys []string
	for key := range d.m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s: %s\n", key, d.m[key])
	}
}
