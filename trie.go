package main

type Node struct {
	Children map[rune]*Node
	Times    int // describe end of word with frequency
}

type Trie struct {
	Root *Node
}

func (t *Trie) Insert(w string, freq int) {
	currentNode := t.Root
	for _, v := range w {
		if currentNode.Children[v] == nil {
			node := make(map[rune]*Node)
			currentNode.Children[v] = &Node{Children: node}
		}
		currentNode = currentNode.Children[v]
	}

	currentNode.Times = freq
}

func (t *Trie) InsertWords(lookUp map[string]int) {
	for k, v := range lookUp {
		t.Insert(k, v)
	}
}

type Suggestion struct {
	Word  string
	Times int
}

var Suggestions []Suggestion

func (n *Node) CollectSuggestions(word string) {
	if n.Times != 0 {
		Suggestions = append(Suggestions, Suggestion{Word: word, Times: n.Times})
	}
	for ch, n := range n.Children {
		n.CollectSuggestions(word + string(ch))
	}
}

func (t *Trie) AutoComplete(userRequest string) {
	// reach the end of key word
	currentNode := t.Root
	for _, v := range userRequest {
		if currentNode.Children[v] == nil {
			return
		}
		currentNode = currentNode.Children[v]
	}

	// collect all words with key prefix
	currentNode.CollectSuggestions(userRequest)
}
