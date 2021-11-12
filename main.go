package main

import (
	"fmt"
	"os"
)

var (
	word      int = 0
	delimiter int = 1
)

type Item struct {
	typ      int
	contents string
}

func main() {
	delimiters := []rune(os.Args[1])
	original := os.Args[2]

	var items []*Item
	isDelimiter := make(map[rune]bool)

	for _, r := range delimiters {
		isDelimiter[r] = true
	}

	var current []rune

	for _, r := range original {
		if isDelimiter[r] {
			if len(current) > 0 {
				newItem := &Item{
					typ:      word,
					contents: string(current),
				}
				items = append(items, newItem)
				current = current[:0]
			}
			newItem := &Item{
				typ:      delimiter,
				contents: string([]rune{r}),
			}
			items = append(items, newItem)
			continue
		}
		current = append(current, r)
	}
	if len(current) > 0 {
		newItem := &Item{
			typ:      word,
			contents: string(current),
		}
		items = append(items, newItem)
	}

	fmt.Println("Original:")
	for _, itm := range items {
		fmt.Printf("%s", itm.contents)
	}
	fmt.Println()

	i, j := 0, len(items)-1

	for i < j {
		if items[i].typ == delimiter {
			i++
			continue
		}
		if items[j].typ == delimiter {
			j--
			continue
		}
		items[i], items[j] = items[j], items[i]
		i++
		j--
	}

	fmt.Println("Words reversed:")
	for _, itm := range items {
		fmt.Printf("%s", itm.contents)
	}
	fmt.Println()
}
