package api

import (
	"crypto/rsa"
	"encoding/pem"
	"ki-kyc/kyc/action"
)

// 获取公钥私钥
func GenerateRSAKey(bits int) (private rsa.PrivateKey, public rsa.PublicKey, err error) {
	return action.GenerateRSAKey(bits)
}

// 获取公钥私钥
func GenerateRSAKeyforPem(bits int) (private pem.Block, public pem.Block, err error) {
	return action.GenerateRSAKeyforPem(bits)
}

// 私钥签名

func SignText(msg []byte, pemtext []byte) (cipher string, err error) {
	return action.SignText(msg, pemtext)
}
