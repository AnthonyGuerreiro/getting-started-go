package custom_sort


func BSort(elements []int) {
	var i int
	var j int
	size := len(elements)

	for i = 0; i < size; i++ {
		for j = i + 1; j < size; j++ {
			if elements[i]>elements[j] {
				swap(elements,i, j)
			}
		}
	}
}