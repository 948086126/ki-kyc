package action

import (
	"crypto/rsa"
)

// sign
type Sign struct {
	Private rsa.PrivateKey `json:"privateKey"`
	Public  rsa.PublicKey  `json:"publicKey"`
	Name    string         `json:"name"`
	Id      string         `json:"id"`
	Type    string         `json:"type"`
}
