package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PriceRequest struct {
	Quantity   int     `json:"quantity"`
	NettoPrice float64 `json:"nettoPrice"`
	Taxes      float64 `json:"taxes"`
}

type PriceResponse struct {
	PriceRequest
	CalculatedBruttoPrice float64 `json:"calculatedBruttoPrice"`
}

func CalculatePrice(c *gin.Context) {
	var priceRequest PriceRequest
	err := c.BindJSON(&priceRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON couldn't be parsed"})
		log.Println("JSON couldn't be parsed")
		log.Println(err)
		return
	}

	price := priceRequest.NettoPrice * float64(priceRequest.Quantity)
	price += price * priceRequest.Taxes

	c.JSON(http.StatusOK, &PriceResponse{
		CalculatedBruttoPrice: price,
		PriceRequest:          priceRequest,
	})
}
