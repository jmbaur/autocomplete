# Autocomplete

An autocompletion engine written in Go.

Uses the trie datastructure to provide possible word completions based on a given dictionary. In the trie, each letter of each word is given a key in a node that is a pointer to the node of the next letter in that word. At the end of a word, a `nil` pointer is used at the key "!".

A server loads a dictionary and listens for fragments of words to complete. A client sends the word fragments. The server/client relationship is implemented with gRPC (very fast 🤓).

```go
type node map[string]*node
```

## Usage

### Server

`./autocompleted <dictionary text file>`

See `assets/dicts` for examples of dictionary files

### Client

`./autocomplete <word fragment>`

### Example

[![asciicast](https://asciinema.org/a/NZ5RbBd9wzY4VKbKQePtUPBL1.svg)](https://asciinema.org/a/NZ5RbBd9wzY4VKbKQePtUPBL1)
