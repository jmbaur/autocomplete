package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmbaur/autocomplete/autocomplete"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: autocomplete <word to complete>")
		os.Exit(1)
	}
	wordToComplete := os.Args[1]

	completer, err := autocomplete.NewCompleter("./dictionary_large.txt")
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
