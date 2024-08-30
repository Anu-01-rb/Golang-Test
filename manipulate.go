package main

import (
	"fmt"
	"sort"
	"strings"
)

func Permutation(text []string, size int, n int, result *[]string) {
	if size == 1 {
		*result = append(*result, strings.Join(text, ""))
		return
	}

	for i := 0; i < size; i++ {
		Permutation(text, size-1, n, result)

		if size%2 == 1 {
			text[0], text[size-1] = text[size-1], text[0]
		} else {
			text[i], text[size-1] = text[size-1], text[i]
		}
	}

}

func RemoveDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Manipulate(text string) []string {
	// TODO : start your code here
	str_split := strings.Split(text, "")
	var result []string
	Permutation(str_split, len(text), len(text), &result)
	result = RemoveDuplicate(result)
	sort.Strings(result)
	return result
}

func main() {
	fmt.Println(Manipulate("ab"))
}
