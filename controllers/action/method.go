package action

import (
	g "github.com/gin-gonic/gin"
	"ki-kyc/kyc/api"
	"log"
	"net/http"
)

//	获取证书
func generateRSAKey(c *g.Context) {
	//(private rsa.PrivateKey, public rsa.PublicKey, err error)
	private, public, err := api.GenerateRSAKey(2048)

	if err != nil {
		log.Println("err:", err)

	}
	Sign := Sign{
		Name:    "sunlidong",
		Id:      "315c6175-cc21-4cd2-b3a9-071d3f57678e",
		Type:    "DHE",
		Private: private,
		Public:  public,
	}

	//
	if err != nil {
		c.JSON(
			http.StatusOK,
			g.H{
				"status": "faild",
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			g.H{"status": "success",
				"data": Sign,
			})
		return
	}

}
