package rand

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"time"
)

func uniqid(prefix string) string {
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000

	return fmt.Sprintf("%s%08x%05x", prefix, sec, usec)
}

func GenerateShortenedURL() string {
	var (
		randomChars   = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
		randIntLength = 27
		stringLength  = 32
	)

	str := make([]rune, stringLength)

	for char := range str {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(randIntLength)))
		if err != nil {
			panic(err)
		}

		str[char] = randomChars[nBig.Int64()]
	}

	hash := sha256.Sum256([]byte(uniqid(string(str))))
	encodedString := base64.URLEncoding.EncodeToString(hash[:])

	return encodedString[0:9]
}
