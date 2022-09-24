package main

import "fmt"

func main() {
	output := [][]string{}
	featurelist := map[string][]string{
		"name":  {"A", "B", "C"},
		"phone": {"1", "2", "3"},
	}
	patterns := []string{"name", "phone", "name"}

	recurse(&output, featurelist, patterns, 0, []string{})

	fmt.Println(output)
}

func recurse(output *[][]string, featurelist map[string][]string, patterns []string, index int, curr []string) {
	if index >= len(patterns) {
		*output = append(*output, curr)
		return
	}

	for _, value := range featurelist[patterns[index]] {
		recurse(output, featurelist, patterns, index+1, append(curr, value))
	}
}
