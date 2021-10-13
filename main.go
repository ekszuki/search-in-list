package main

import (
	"fmt"
	"strings"
)

func main() {
	target := "b-9"

	wordList := genValues(100)

	find := createFindFunction(50)
	retList := find(wordList, target)

	fmt.Println(retList)
}

func createFindFunction(retSize int) func(wordList []string, target string) []string {
	count := 0
	return func(wordList []string, target string) []string {
		retList := []string{}
		for _, w := range wordList {
			if len(w) < len(target) {
				continue
			}

			pw := w[0:len(target)]
			if !strings.EqualFold(pw, target) {
				continue
			}

			retList = append(retList, w)
			count++
			if count == retSize {
				return retList
			}
		}

		return retList
	}
}

func genValues(size int) []string {
	var wordList []string
	for i := 0; i < size; i++ {
		w := ""
		if i%2 == 0 {
			w = fmt.Sprintf("a-%d", i)
		} else {
			w = fmt.Sprintf("b-%d", i)
		}
		wordList = append(wordList, w)
	}

	return wordList
}
