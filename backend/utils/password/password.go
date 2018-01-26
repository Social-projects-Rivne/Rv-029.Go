package password

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
)

// encrypt password with salt
func EncodePassword(md5Hash string, salt string) string {
	var buffer bytes.Buffer

	buffer.WriteString(md5Hash)
	buffer.WriteString(salt)

	hasher := sha256.New()
	hasher.Write(buffer.Bytes())

	return hex.EncodeToString(hasher.Sum(nil))
}