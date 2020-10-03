package autocomplete

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Completer struct {
	Path       string
	Dictionary trie
}

/* NewCompleter takes in the path of the dictionary to use for completion */
func NewCompleter(path string) (*Completer, error) {
	fullpath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	newCompleter := Completer{Path: fullpath}
	err = newCompleter.load()
	if err != nil {
		return nil, err
	}
	return &newCompleter, nil
}

func (c *Completer) load() error {
	data, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return err
	}

	c.Dictionary = trie{base: &node{}}
	word := ""
	for _, r := range string(data) {
		if string(r) == "\n" {
			c.Dictionary.insert(word)
			word = ""
		} else {
			word += string(r)
		}
	}
	return nil
}

func (c *Completer) Complete(fragment string) ([]string, error) {
	existsCompletions := c.Dictionary.find(fragment)
	if !existsCompletions {
		fmt.Println("Not a word")
	}
	return []string{}, nil
}
