package api

import (
	"crypto/rsa"
	"ki-kyc/kyc/action"
)

// 获取公钥私钥
func GenerateRSAKey(bits int) (private rsa.PrivateKey, public rsa.PublicKey, err error) {
	return action.GenerateRSAKey(bits)
}
