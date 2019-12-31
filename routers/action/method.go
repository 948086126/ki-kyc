package action

import (
	g "github.com/gin-gonic/gin"
	Kiapi "ki-kyc/controllers/api"
)

func initRouter() *g.Engine {

	router := g.Default()

	api := router.Group("/api")

	v1 := api.Group("/v1")

	Test := v1.Group("/test")

	{
		Test.POST("/getCert", Kiapi.GenerateRSAKey)
		Test.POST("/getCert2", Kiapi.GenerateRSAKey)
		Test.POST("/getCert3", Kiapi.GenerateRSAKey)
	}

	//
	return router
}
