package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomString generates a random string of the specified length
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomSeed := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSeed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[randomGenerator.Intn(len(charset))]
	}
	return string(result)
}
