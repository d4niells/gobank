package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqestuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min+1)
}

// RandomStr generates a random string of length n
func RandomStr(n int) string {
	var str strings.Builder
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(26)]
		str.WriteByte(c)
	}

	return str.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomStr(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generate a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	return currencies[rand.Intn(len(currencies))]
}