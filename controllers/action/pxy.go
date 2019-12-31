package action

import (
	g "github.com/gin-gonic/gin"
)

// 获取公钥私钥
func GenerateRSAKey(c *g.Context) {
	generateRSAKey(c)
}
