package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijlkmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder // Узнать получше,  как это работает
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	
	return sb.String()
}

func RandomOwner() string {
	return RandomString(7)
}

func RandomMoney(from, to int64) int64 {
	return RandomInt(from, to)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "RUB"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}