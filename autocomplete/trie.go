package autocomplete

import "fmt"

const end string = "!"

type node map[string]*node
type trie struct {
	base *node
}

func (t trie) insert(word string) {
	current := *t.base
	for _, r := range word {
		if current[string(r)] == nil {
			current[string(r)] = &node{}
		}
		current = *current[string(r)]
	}
	current[end] = nil
}

func (t trie) find(word string) bool {
	if t.base == nil {
		return false
	}
	current := *t.base
	for _, r := range word {
		fmt.Println(current)
		_, ok := current[string(r)]
		// if word is not found
		if !ok {
			return false
		}
		current = *current[string(r)]
	}
	// if at end of word
	if _, atEnd := current[end]; atEnd {
		return true
	}
	return false
}
