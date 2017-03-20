package prime

import "math"

func NextPrime(n int64) int64 {
	var sqrt = int32(math.Sqrt(float64(n)))

	var i int32

	for i = 2; i < sqrt; i++ {
		var factors = get_factors(i, n)
		if len(factors) != 0 {
			return NextPrime(n + 1)
		}
	}
	return n
}
