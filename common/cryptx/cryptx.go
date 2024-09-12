package cryptx

/**
 * @Author elastic·H
 * @Date 2024-09-12
 * @File: cryptx.go
 * @Description:
 */

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}
