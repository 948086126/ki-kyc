package action

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"log"
)

//	获取rsa私钥  参数：int
func generateRSAKey(bits int) (private rsa.PrivateKey, public rsa.PublicKey, err error) {

	if bits == 0 {
		log.Println("bits is nil")
		return rsa.PrivateKey{}, rsa.PublicKey{}, errors.New("bits is nil")
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		log.Println("generateRSAKey err:", err)
		return rsa.PrivateKey{}, rsa.PublicKey{}, err
	}

	log.Println("generateRSAKey ok:", privateKey.Size())

	return *privateKey, privateKey.PublicKey, err
}
