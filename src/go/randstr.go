package randstr

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, charsetLen)
		result[i] = charset[num.Int64()]
	}

	return string(result)
}
