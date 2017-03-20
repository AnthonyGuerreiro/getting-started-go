package coin

import "math/rand"

//returns true if flips heads, false if tails
func Flip() bool {
	return rand.Int()%2 == 0
}

func IsHeads(bool bool) bool {
	return bool
}

func IsTails(bool bool) bool {
	return !IsHeads(bool)
}
