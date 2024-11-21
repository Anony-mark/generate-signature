package controller

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type infoBody struct {
	MerchantCode string `bson:"merchantCode" json:"merchantCode" binding:"required"`
	MerchantKey  string `bson:"merchantKey" json:"merchantKey" binding:"required"`
	Currency     string `bson:"currency" json:"currency" binding:"required"`
	PaymentID    string `bson:"paymentID" json:"paymentID" binding:"required"`
	ResponseURL  string `bson:"responseURL" json:"responseURL" binding:"required"`
	Amount       string `bson:"amount" json:"amount" binding:"required"`
}

func GenerateSignature() func(c *gin.Context) {
	return func(c *gin.Context) {
		bodyInfo := infoBody{}

		if err := c.ShouldBindJSON(&bodyInfo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// fmt.Println("bodyInfo:", bodyInfo)

		data := url.Values{}
		data.Set("merchantCode", bodyInfo.MerchantCode)
		data.Set("merchantKey", bodyInfo.MerchantKey)
		data.Set("currency", bodyInfo.Currency)
		data.Set("paymentID", bodyInfo.PaymentID)
		data.Set("responseURL", bodyInfo.ResponseURL)
		data.Set("amount", bodyInfo.Amount)

		reqBody := strings.NewReader(data.Encode())

		// fmt.Println("Request Body:", reqBody)

		resp, err := http.Post("https://demo.zpay.global/signature", "application/x-www-form-urlencoded", reqBody)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Println("Error reading response body:", err)
		// 	return
		// }

		body, _ := io.ReadAll(resp.Body)

		// fmt.Println("response Status:", string(body))

		// โหลด HTML ด้วย goquery
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
		if err != nil {
			log.Fatal(err)
		}

		// ดึงค่าจาก div ที่มี class "alert alert-success"
		signature := strings.TrimSpace(strings.Replace(doc.Find(".alert.alert-success").Text(), "Generated Signature:", "", 1))

		fmt.Println("Generated Signature:", signature)

		c.JSON(200, gin.H{"message": "Success", "data": signature})
	}
}
