package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"paynicornDemo/demo"
)



type PostbackReq struct{
	Content string `json:"content"`
	Sign string `json:"sign"`
}

func main() {

	r := gin.Default()

	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)


	r.GET("/pay", func(context *gin.Context){
		countryCode := context.Query("countryCode")
		orderId := context.Query("orderId")
		orderDescription := context.Query("orderDescription")
		currency := context.Query("currency")
		amount := context.Query("amount")
		appKey := context.Query("appkey")
		merchantSecret := context.Query("merchantsecret")

		context.JSON(http.StatusOK,demo.PaymentDemo(countryCode,orderId,orderDescription,currency,amount,appKey,merchantSecret))

	})

	r.POST("/postback", func(context *gin.Context) {
		var req PostbackReq
		if err := context.BindJSON(&req); err != nil{
			fmt.Println("bind json failed")
		}else{
			p := demo.PostbackDemo(req.Content,req.Sign)
			if p != nil{
				fmt.Println("succeed :"+p.TxnId)
				context.String(http.StatusOK, "success_"+p.TxnId)
			}
		}
	})

	r.GET("/query", func(context *gin.Context){

		orderId := context.Query("orderId")
		txnType := context.Query("txnType")
		appKey := context.Query("appkey")
		merchantSecret := context.Query("merchantsecret")

		context.JSON(http.StatusOK,demo.QueryDemo(orderId,txnType,appKey,merchantSecret))

	})



	// Listen and Server in 0.0.0.0:8080
	r.Run(":80")
}
