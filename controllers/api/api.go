package api

import (
	g "github.com/gin-gonic/gin"
	"ki-kyc/controllers/action"
)

// 获取公钥私钥
func GenerateRSAKey(c *g.Context) {
	action.GenerateRSAKey(c)
}

func GenerateRSAKeyforPem(c *g.Context) {
	action.GenerateRSAKeyforPem(c)
}

//	数字签名
func Encryption(c *g.Context) {
	action.Encryption(c)
}
