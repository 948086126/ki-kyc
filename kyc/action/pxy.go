package action

import (
	"crypto/rsa"
	"encoding/pem"
)

// 获取公钥私钥
func GenerateRSAKey(bits int) (private rsa.PrivateKey, public rsa.PublicKey, err error) {
	return generateRSAKey(bits)
}

// 获取公钥私钥
func GenerateRSAKeyforPem(bits int) (private pem.Block, public pem.Block, err error) {
	return generateRSAKeyforPem(bits)
}

// 私钥签名
func SignText(msg []byte, pemtext []byte) (cipher string, err error) {
	return signText(msg, pemtext)
}
