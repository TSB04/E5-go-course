package main

import "fmt"

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) {
    d[word] = definition
}

func (d Dictionary) Get(word string) (string, bool) {
    definition, exists := d[word]
    return definition, exists
}

func (d Dictionary) Remove(word string) {
    delete(d, word)
}

func (d Dictionary) List() {
    if len(d) == 0 {
        fmt.Println("The dictionary is empty.")
        return
    }
    fmt.Println("Dictionary contents:")
    for word, definition := range d {
        fmt.Printf("%s: %s\n", word, definition)
    }
}

func main() {
    dico := Dictionary{}

    // Adding words
    dico.Add("Go", "A statically typed, compiled programming language.")
    dico.Add("Goroutine", "A lightweight thread managed by the Go runtime.")
    dico.Add("Channel", "A communication mechanism between Goroutines.")

    // Listing words
    dico.List()

    // Getting a definition
    word := "Go"
    if definition, found := dico.Get(word); found {
        fmt.Printf("\nDefinition of %s: %s\n", word, definition)
    } else {
        fmt.Printf("\n%s not found in dictionary.\n", word)
    }

    // Removing a word
    dico.Remove("Goroutine")

    // Listing after removal
    fmt.Println("\nAfter removal:")
    dico.List()
}
