package main

import (
	"fmt"

	"github.com/lucasres/db-stack/pkg"
)

var (
	m *pkg.Manager
)

func getManager() *pkg.Manager {
	if m == nil {
		m = pkg.NewManager()
	}

	return m
}

func main() {
	// create a singletron for manager
	manager := getManager()
	// set a new key value
	manager.Set("key", "val1")
	// print key value
	fmt.Println(manager.Get("key"))
	// push a new value
	manager.Push("key", "val2")
	// print key value
	fmt.Println(manager.Get("key"))
	// pop value
	manager.Pop("key")
	// print a key
	fmt.Println(manager.Get("key"))
	// clean database
	manager.Flush()
	// try get a key
	fmt.Println(manager.Get("key"))
}
