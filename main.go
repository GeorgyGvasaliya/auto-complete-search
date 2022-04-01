package main

import (
	"fmt"
	"sort"
)

func PrintSuggestions(sl []Suggestion) {
	if len(Suggestions) > 5 {
		sl = sl[:5]
	}
	for _, v := range sl {
		fmt.Printf("%v\n", v.Word)
	}
}

func main() {
	// набираем п в инкогнито яндекс браузере
	//inputData := map[string]int{
	//	"переводчик":                          12,
	//	"почта россии отслеживание":           2,
	//	"переводчик с английского на русский": 1,
	//	"почта": 4,
	//	"программа передач на сегодня": 9,
	//	"погода":          7,
	//	"погода в москве": 8,
	//	"производственный календарь 2022": 5,
	//	"почта маил": 6,
	//}
	inputData := map[string]int{
		"переводчик": 1,
	}
	// init Trie
	rootChildren := make(map[rune]*Node)
	rootNode := &Node{Children: rootChildren}
	t := Trie{Root: rootNode}

	// insertionss
	t.InsertWords(inputData)

	// driver func
	t.AutoComplete("переводчик")
	//t.AutoComplete("vqihfoehqc") // empty output

	// ranking results
	sort.Slice(Suggestions, func(i, j int) bool {
		return Suggestions[i].Times > Suggestions[j].Times
	})

	// print answer
	PrintSuggestions(Suggestions)

}
