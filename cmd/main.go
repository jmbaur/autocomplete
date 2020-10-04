package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmbaur/autocomplete/autocomplete"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: autocomplete <word to complete>")
		os.Exit(1)
	}
	dictionary := os.Args[1]
	wordToComplete := os.Args[2]

	completer, err := autocomplete.NewCompleter(dictionary)
	if err != nil {
		log.Fatal(err)
	}
	possibleWords, err := completer.Complete(wordToComplete)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range possibleWords {
		fmt.Println(v)
	}
}
