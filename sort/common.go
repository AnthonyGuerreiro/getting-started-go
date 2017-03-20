package custom_sort

import "math/rand"

func swap(elements []int, i int, j int) {
	elements[i], elements[j] = elements[j], elements[i]
}

func getRandomElementBetween(elements []int, start int, end int) int {
	var size = end - start
	return elements[start+rand.Int()%size]
}
