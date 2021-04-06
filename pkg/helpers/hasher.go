package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

type PasswordHasher interface {
	PasswordToMD5Hash(pass string) string
}

type Md5 struct {
	Salt string
}

func NewHasher(salt string) *Md5 {
	return &Md5{
		Salt: salt,
	}
}

func (m *Md5) PasswordToMD5Hash(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))

	return hex.EncodeToString(append(hasher.Sum(nil), []byte(m.Salt)...))
}
