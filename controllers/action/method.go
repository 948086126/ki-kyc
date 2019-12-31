package action

import (
	"encoding/json"
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
	ByteSign, err := json.Marshal(&Sign)
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
				"data": string(ByteSign),
			})
		return
	}

}

//	获取证书
func generateRSAKeyforPem(c *g.Context) {
	//(private rsa.PrivateKey, public rsa.PublicKey, err error)
	private, public, err := api.GenerateRSAKeyforPem(2048)

	if err != nil {
		log.Println("err:", err)

	}
	Sign := SignPem{
		Name:    "sunlidong",
		Id:      "315c6175-cc21-4cd2-b3a9-071d3f57678e",
		Type:    "DHE",
		Private: private,
		Public:  public,
	}
	log.Println("Sign", Sign)
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

//  加密
func encryption(c *g.Context) {
	Pg := SignText{}
	if err := c.ShouldBindJSON(&Pg); err != nil {
		log.Println("err2:", err)
	}
	if Pg.Msg == "" || Pg.Public.Type == "" || Pg.Private.Type == "" {
		log.Println("2data is nil ")
	}

	// 获取数据

	// 校验数据

	// 数据加密
	result, err := api.SignText([]byte(Pg.Msg), Pg.Private.Bytes)

	//	判断结果
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
				"data": string(result),
			})
		return
	}
}
