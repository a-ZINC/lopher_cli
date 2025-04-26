package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"slices"
)

func Hash(alg string, data []byte) string {
	switch alg {
	case "sha256":
		buffer := sha256.Sum256(data)
		return hex.EncodeToString(buffer[:])
	case "sha512":
		buffer := sha512.Sum512(data)
		return hex.EncodeToString(buffer[:])
	case "md5":
		buffer := md5.Sum(data)
		return hex.EncodeToString(buffer[:])
	default:
		return "sha256"
	}
}

func Contains(s []string, e string) bool {
	return slices.Contains(s, e)
}
