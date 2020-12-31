package demo

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


var MERCHANT_SCRECT = "PUT_YOUR_MERCHANT_MD5_SECRET_KEY_HERE"
var APP_KEY = "PUT_YOUR_APP_KEY_HERE"




type Content struct{
	Content map[string]string
}

func (c*Content)AddContent(key string,value string){
	if c.Content == nil{
		c.Content = make(map[string]string)
	}

	c.Content[key] = value
}

func (c*Content)GetContentBase64()string{
	str,err :=json.Marshal(c.Content)

	if err == nil{
		encode := base64.StdEncoding.EncodeToString(str)
		return encode
	}

	return ""
}

type RequestBody struct{
	Content string `json:"content"`
	Sign string `json:"sign"`
	AppKey string `json:"appKey"`
}

type ResponseBody struct{
	ResponseCode string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Content string `json:"content"`
	Sign string `json:"sign"`
}




func PaymentDemo(countryCode string,orderId string,orderDescription string,currency string,amount string)string{

	c := Content{}

	c.AddContent("countryCode",countryCode)
	c.AddContent("orderId",orderId)
	c.AddContent("orderDescription",orderDescription)
	c.AddContent("currency",currency)
	c.AddContent("amount",amount)

	b := RequestBody{}

	b.Content = c.GetContentBase64()
	b.Sign = fmt.Sprintf("%x",md5.Sum([]byte(b.Content+MERCHANT_SCRECT)))
	b.AppKey = APP_KEY



	client := &http.Client{}

	jsonBytes, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	request, err := http.NewRequest("POST", "https://api.paynicorn.com/trade/v3/transaction/pay", strings.NewReader(string(jsonBytes)))
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")

	var buffer []byte
	if response, err := client.Do(request); err != nil {

	} else {
		if buffer, err = ioutil.ReadAll(response.Body); err != nil {
			fmt.Println(err)
		}
	}

	rsp := ResponseBody{}
	err = json.Unmarshal(buffer, &rsp)


	if rsp.ResponseCode == "000000"{

		if sign := fmt.Sprintf("%x",md5.Sum([]byte(rsp.Content+MERCHANT_SCRECT))); sign == rsp.Sign{

			content, err := base64.StdEncoding.DecodeString(rsp.Content)
			if err == nil {
				return string(content)

			}
		}
	}else{
		fmt.Println(rsp.ResponseMessage)

	}

	return ""
}


type Postback struct{
	TxnId string `json:"txnId"`
	OrderId string `json:"orderId"`
	Amount string `json:"amount"`
	Currency string `json:"currency"`
	CountryCode string `json:"countryCode"`
	Status string `json:"status"`
	Code string `json:"code"`
	Message string `json:"message"`
}

func PostbackDemo(content string,sign string)*Postback{
	if s := fmt.Sprintf("%x",md5.Sum([]byte(content+MERCHANT_SCRECT))); sign == s{

		c, err := base64.StdEncoding.DecodeString(content)
		if err == nil {
			p :=&Postback{}
			err = json.Unmarshal(c,p)
			if err == nil{
				return p
			}
		}
	}else{
		fmt.Println("sign verify failed")
	}

	return nil
}


func QueryDemo(orderId string,txnType string)string{

	c := Content{}
	c.AddContent("orderId",orderId)
	c.AddContent("txnType",txnType)

	b := RequestBody{}

	b.Content = c.GetContentBase64()
	b.Sign = fmt.Sprintf("%x",md5.Sum([]byte(b.Content+MERCHANT_SCRECT)))
	b.AppKey = APP_KEY



	client := &http.Client{}

	jsonBytes, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	request, err := http.NewRequest("POST", "https://api.paynicorn.com/trade/v3/transaction/query", strings.NewReader(string(jsonBytes)))
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")

	var buffer []byte
	if response, err := client.Do(request); err != nil {

	} else {
		if buffer, err = ioutil.ReadAll(response.Body); err != nil {
			fmt.Println(err)
		}
	}

	rsp := ResponseBody{}
	err = json.Unmarshal(buffer, &rsp)


	if rsp.ResponseCode == "000000"{

		if sign := fmt.Sprintf("%x",md5.Sum([]byte(rsp.Content+MERCHANT_SCRECT))); sign == rsp.Sign{

			content, err := base64.StdEncoding.DecodeString(rsp.Content)
			if err == nil {
				return string(content)

			}
		}
	}else{
		fmt.Println(rsp.ResponseMessage)

	}

	return ""

}

func RefundDemo(){

}

func AuthpayDemo(){

}

func PayoutDemo(){

}
