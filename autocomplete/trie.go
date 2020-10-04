package autocomplete

import (
	"sort"
)

const end string = "!"

var words []string

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

func (t trie) find(word string) []string {
	if t.base == nil {
		return []string{}
	}
	current := *t.base
	for _, r := range word {
		_, ok := current[string(r)]
		// if word is not found
		if !ok {
			return []string{}
		}
		current = *current[string(r)]
	}

	traverse(current, word)
	sort.SliceStable(words, func(i, j int) bool {
		if len(words[i]) > len(words[j]) {
			return false
		}
		return true
	})
	return words
}

func traverse(curNode node, word string) {
	for k, next := range curNode {
		if k == end {
			words = append(words, word)
		} else {
			traverse(*next, word+k)
		}
	}
}
