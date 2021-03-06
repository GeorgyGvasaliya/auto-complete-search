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
	inputData := map[string]uint16{
		"переводчик":                          12,
		"почта россии отслеживание":           2,
		"переводчик с английского на русский": 1,
		"почта":                               4,
		"программа передач на сегодня":        9,
		"погода":                              7,
		"погода в москве":                     8,
		"производственный календарь 2022":     5,
		"почта маил":                          6,
	}
	//inputData := map[string]uint16{
	//	"league of legends": 3,
	//	"learn javascript":  7,
	//	"leaf":              1,
	//	"leonardo dicaprio": 10,
	//	"leonardo da vinci": 8,
	//}
	//inputData := map[string]uint16{
	//	"how to search by image":          6,
	//	"how to search for pdf on google": 12,
	//	"how to search in telegram":       10,
	//	"how to search in google":         20,
	//	"how to search in google sheets":  9,
	//	"how to search in google docs":    7,
	//}
	//inputData := map[string]uint16{
	//	"подарить машину тикток":   3,
	//	"подари жизнь":             5,
	//	"подарки на 23 февряля":    10,
	//	"подарочный сертификат":    3,
	//	"подарки мужчине":          4,
	//	"подарки на день рождения": 5,
	//	"подагра":                  1,
	//	"подать заявление в загс":  1,
	//}

	userRequest := "перев"
	//userRequest := "iuqdieqndiwqndiewqbfc" // -> no hints

	// init Trie
	rootChildren := make(map[rune]*Node)
	rootNode := &Node{Children: rootChildren}
	t := Trie{Root: rootNode}

	// insertions
	t.InsertWords(inputData)

	// driver func
	t.AutoComplete(userRequest)

	// ranking results
	sort.Slice(Suggestions, func(i, j int) bool {
		return Suggestions[i].Times > Suggestions[j].Times
	})

	// print answer
	PrintSuggestions(Suggestions)
}
