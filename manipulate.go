package main

import (
	"strings"
)

func Permutation(text []string, size int, n int) []string {
	permuted_list := make([]string, 0)
	if size == 1 {
		return text
	}

	for i := 0; i < size; i++ {
		arranged := Permutation(text, size-1, n)

		str_arranged := strings.Join(arranged, "")
		if size == n {
			permuted_list = append(permuted_list, arranged...)
		} else {
			permuted_list = append(permuted_list, str_arranged)
		}

		if size%2 == 1 {
			text[0], text[size-1] = text[size-1], text[0]
		} else {
			text[i], text[size-1] = text[size-1], text[i]
		}
	}

	return permuted_list
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

func Manipulate(text []string) []string {
	// TODO : start your code here
	man := Permutation(text, len(text), len(text))
	man = RemoveDuplicate(man)
	return man
}
