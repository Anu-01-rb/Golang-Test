package main

func FindOddNumber(text []int) int {
	// TODO : start your code here
	key := make(map[int]int)

	for _, num := range text {
		key[num]++
	}

	for num, count := range key {
		if count%2 == 1 {
			return num
		}
	}

	return 0
}
