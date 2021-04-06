package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

type PasswordHasher interface {
	NewMD5Hash() (string, error)
}

type Md5 struct {
	Salt string
}

func NewHasher(salt string) *Md5 {
	return &Md5{
		Salt: salt,
	}
}

func (m *Md5) NewMD5Hash(pass string) (string, error) {
	hasher := md5.New()

	_, err := hasher.Write([]byte(pass))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(append(hasher.Sum(nil), []byte(m.Salt)...)), nil
}
