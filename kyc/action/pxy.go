package action

import (
	"crypto/rsa"
)

// 获取公钥私钥
func GenerateRSAKey(bits int) (private rsa.PrivateKey, public rsa.PublicKey, err error) {
	return generateRSAKey(bits)
}
