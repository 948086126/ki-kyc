package action

import (
	g "github.com/gin-gonic/gin"
)

// 获取公钥私钥
func GenerateRSAKey(c *g.Context) {
	generateRSAKey(c)
}

// 获取公钥私钥
func GenerateRSAKeyforPem(c *g.Context) {
	generateRSAKeyforPem(c)
}

// encryption
// 获取公钥私钥
func Encryption(c *g.Context) {
	encryption(c)
}
