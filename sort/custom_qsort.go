package custom_sort


func Qsort(elements []int) {
	qsort_range(elements, 0, len(elements))
}

func qsort_range(elements []int, start int, end int) {

	if end-start < 2 {
		return
	}

	var pivot = getRandomElementBetween(elements, start, end)

	var i = 0
	var j = end - 1

	for ; i < j; {
		if elements[i] < pivot {
			i++
			continue
		}
		if elements[j] > pivot {
			j--
			continue
		}
		swap(elements, i, j)
		i++
		j--
	}
	qsort_range(elements, start, i)
	qsort_range(elements, i, end)
}
