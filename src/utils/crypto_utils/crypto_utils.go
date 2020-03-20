package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
)

func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateSecret(length int) string {
	data := make([]byte, length)
	_, _ = rand.Read(data)
	return fmt.Sprintf("%x", data)
}