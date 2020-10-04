# Autocomplete

An autocompletion engine written in Go.

Uses the trie datastructure to provide possible word completions based on a given dictionary. In the trie, each letter of each word is given a key in a node that is a pointer to the node of the next letter in that word. At the end of a word, a `nil` pointer is used at the key "!".

```go
type node map[string]*node
```

## Usage

Input:
`./autocomplete waterm`

Output:

```
waterman
watermanship
watermark
watermelon
```

## TODO

- [ ] traverse all possible words after the last node is found for a given word fragment
