package main

import (
	"fmt"
	"sort"
)

var m map[string]int

func get(key string) (int, bool) {
	value, found := m[key]
	return value, found
}

func add(key string, value int) {
	m[key] = value
}

func remove(key string) {
	delete(m, key)
}

func list(m map[string]int) {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, m[key])
	}
}

func main() {
	m = make(map[string]int)

	add("rr", 1)
	add("ll", 2)
	add("dd", 3)

	if value, found := get("ff"); found {
		fmt.Printf("Value found: %d\n", value)
	} else {
		fmt.Println("Key not found.")
	}
	fmt.Println("Avant remove", m)
	remove("ll")
	fmt.Println("Apres remove", m)

	list(m)
}
