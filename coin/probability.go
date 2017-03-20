package coin

func FlipNTimes(times int) (int) {
	var i int
	var heads = 0
	for i = 0; i < times; i++ {
		if IsHeads(Flip()) {
			heads++
		}
	}
	return heads
}
