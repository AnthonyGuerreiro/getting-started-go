package fibonacci

func Dyn(index int) int64 {

	var previous int64 = 1
	var current int64 = 1

	var i int
	for i = 2; i < index; i++ {
		var next = previous + current
		previous = current
		current = next
	}
	return current
}

func Rec(index int) int64 {
	if index == 1 || index == 2 {
		return 1
	}
	return Rec(index-1) + Rec(index-2)
}
