package main

import (
	"ki-kyc/routers/api"
	"log"
)

func main() {
	Eng := api.InitRouter()
	Eng.Run(":10081")
	log.Println("服务启动成功")
}
