package main

type Node struct {
	Children map[rune]*Node
	Times    uint16 // describe end of word with frequency
}

type Trie struct {
	Root *Node
}

func (t *Trie) Insert(word string, freq uint16) {
	currentNode := t.Root
	for _, v := range word {
		if currentNode.Children[v] == nil {
			node := make(map[rune]*Node)
			currentNode.Children[v] = &Node{Children: node}
		}
		currentNode = currentNode.Children[v]
	}

	currentNode.Times = freq
}

func (t *Trie) InsertWords(inputData map[string]uint16) {
	for k, v := range inputData {
		t.Insert(k, v)
	}
}

type Suggestion struct {
	Word  string
	Times uint16
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
		if currentNode.Children[v] == nil { // no such prefix
			return
		}
		currentNode = currentNode.Children[v]
	}

	// collect all words with key prefix
	currentNode.CollectSuggestions(userRequest)
}
