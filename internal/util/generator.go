package util

import "math/rand"

func GenerateRandomString(n int) string {
	var charsets = []rune("ABCDEFGHIJabcdefghij12345678910")
	letters := make([]rune, n)
	for i := range letters {
		letters[i] = charsets[rand.Intn(len(charsets))]
	}
	return string(letters)
}
