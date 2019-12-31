package api

import (
	g "github.com/gin-gonic/gin"
	"ki-kyc/controllers/action"
)

// 获取公钥私钥
func GenerateRSAKey(c *g.Context) {
	action.GenerateRSAKey(c)
}
