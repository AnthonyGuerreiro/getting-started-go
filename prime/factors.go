package prime

func Factors(n int64) []int32 {
	if n == 1 {
		return []int32{1}
	}
	var factors []int32

	var i int32

	for i = 2; n != 1; i++ {
		var factors_i = get_factors(i, n)
		n = divide(n, factors_i)
		factors = append(factors, factors_i...)
	}

	return factors
}

func divide(n int64, factors []int32) int64 {
	for _, factor := range factors {
		n /= int64(factor)
	}
	return n
}
