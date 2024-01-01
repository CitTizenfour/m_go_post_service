package helper

import (
	"math"
	"math/rand"
	"time"
)

const POOL = "1234567890abcdefghijklmnopqrstuwxvyzABcDEFGHIJKLMNOPQRSTUYVWXYZ"

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func RoundFloat2DecimalPrecison(f float64) float64 {
	return math.Floor(f*100) / 100
}

func RoundFloat4DecimalPrecison(f float64) float64 {
	return math.Floor(f*10000) / 10000
}

func GenerateRandomString(n int) string {
	l := byte(len(POOL))

	b, err := GenerateRandomBytes(n)
	if err != nil {
		return ""
	}

	for i := 0; i < n; i++ {
		b[i] = POOL[(b[i])%l]
	}

	return string(b)
}

// GenerateRandomBytes returns securely generated random bytes
func GenerateRandomBytes(n int) ([]byte, error) {
	rand.Seed(time.Now().Unix())
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
