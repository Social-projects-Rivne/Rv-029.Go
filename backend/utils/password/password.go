package password

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// encrypt password with salt
func EncodePassword(md5Hash string, salt string) string {
	var buffer bytes.Buffer

	buffer.WriteString(md5Hash)
	buffer.WriteString(salt)

	hasher := sha256.New()
	hasher.Write(buffer.Bytes())

	return hex.EncodeToString(hasher.Sum(nil))
}

// generate random salt string
func GenerateSalt(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func EncodeMD5(string string) string {
	hasher := md5.New()
	hasher.Write([]byte(string))

	return hex.EncodeToString(hasher.Sum(nil))
}
