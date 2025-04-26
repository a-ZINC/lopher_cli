package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"slices"
)

func Hash(alg string, data string) string {
	switch alg {
	case "sha256":
		buffer := sha256.Sum256([]byte(data))
		return string(buffer[:])
	case "sha512":
		buffer := sha512.Sum512([]byte(data))
		return string(buffer[:])
	case "md5":
		buffer := md5.Sum([]byte(data))
		return string(buffer[:])
	default:
		return "sha256"
	}
}

func Contains(s []string, e string) bool {
	return slices.Contains(s, e)
}
