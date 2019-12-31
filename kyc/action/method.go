package action

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
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
func generateRSAKeyforPem(bits int) (private pem.Block, public pem.Block, err error) {

	if bits == 0 {
		log.Println("bits is nil")
		return pem.Block{}, pem.Block{}, errors.New("bits is nil")
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		log.Println("generateRSAKey err:", err)
		return pem.Block{}, pem.Block{}, err
	}

	log.Println("generateRSAKey ok:", privateKey.Size())

	//
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := pem.Block{
		Type:  "RSA Private Key",
		Bytes: X509PrivateKey,
	}
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	publicBlock := pem.Block{
		Type:  "RSA Public Key",
		Bytes: X509PublicKey,
	}

	//
	privateFile, err := os.Create("private.pem")
	if err != nil {
		log.Println("err:", err)
		return pem.Block{}, pem.Block{}, nil
	}
	publicFile, err := os.Create("public.pem")
	if err != nil {
		log.Println("err:", err)
		return pem.Block{}, pem.Block{}, nil
	}
	defer privateFile.Close()
	defer publicFile.Close()

	//
	pem.Encode(privateFile, &privateBlock)
	pem.Encode(publicFile, &publicBlock)

	return privateBlock, publicBlock, err
}

//	签名
func signText(msg []byte, pemtext []byte) (cipher string, err error) {

	//	解码
	privateKey, err := x509.ParsePKCS1PrivateKey(pemtext)
	if err != nil {
		log.Println("解码 5err:", err)
		return "", nil
	}
	//计算散列值
	hash := sha256.New()
	hash.Write(msg)
	bytes := hash.Sum(nil)
	// 签名
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, bytes)
	if err != nil {
		log.Println("err6:", err)
		return "", nil
	}

	//
	log.Println("msg:", string(sign))
	return string(sign), nil
}
