package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func NewMD5Hash(pass string) (string, error) {
	hasher := md5.New()

	_, err := hasher.Write([]byte(pass))
	if err != nil {
		return "", errors.New("string hasher error")
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
