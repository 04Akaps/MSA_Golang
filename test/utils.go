package test

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func randomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomName() string {
	return randomString(6)
}

func RandomAge() int64 {
	return randomInt(0, 100)
}
