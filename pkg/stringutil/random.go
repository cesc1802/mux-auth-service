package stringutil

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// len = 48
// random(98932849) % len = [0, len-1]
func randomStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(98932849)%len(letters)]
	}
	return string(b)
}

func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}
	return randomStr(length)
}
