package api

import (
	g "github.com/gin-gonic/gin"
	"ki-kyc/routers/action"
)

//
func InitRouter() *g.Engine {
	return action.InitRouter()
}
