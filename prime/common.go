package prime

func get_factors(factor int32, n int64) []int32 {
	var factors []int32

	for ; n%int64(factor) == 0; n /= int64(factor) {
		factors = append(factors, factor)
	}
	return factors
}