package action

import (
	"crypto/rsa"
	"encoding/pem"
)

// sign
type Sign struct {
	Private rsa.PrivateKey `json:"privateKey"`
	Public  rsa.PublicKey  `json:"publicKey"`
	Name    string         `json:"name"`
	Id      string         `json:"id"`
	Type    string         `json:"type"`
}

type SignPem struct {
	Private pem.Block `json:"privateKey"`
	Public  pem.Block `json:"publicKey"`
	Name    string    `json:"name"`
	Id      string    `json:"id"`
	Type    string    `json:"type"`
}

// 数据加密

type SignText struct {
	Private pem.Block `json:"privateKey"`
	Public  pem.Block `json:"publicKey"`
	Msg     string    `json:"msg"`
}
