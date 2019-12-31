package main

import (
	"ki-kyc/routers/api"
)

func main() {
	Eng := api.InitRouter()
	Eng.Run(":10081")
}
