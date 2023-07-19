package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	result := map[string]int{}

	var label string

	for _, x := range s {
		if x == ' ' {
			if label != "" {
				result[label]++
			}
			label = ""
		} else {
			label += string(x)
		}
	}

	if label != "" {
		result[label]++
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
