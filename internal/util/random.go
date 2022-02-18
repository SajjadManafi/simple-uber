package util

import (
	"fmt"
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random int in the range [min, max]
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string of length n
func RandomString(length int) string {
	k := len(alphabet)
	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[rand.Intn(k)]
	}
	return string(b)
}

// RandomUsername generates a random username
func RandomUsername() string {
	return RandomString(6)
}

// // RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
