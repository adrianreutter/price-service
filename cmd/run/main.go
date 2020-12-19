package main

import (
	"log"
	"time"

	"github.com/adrianreutter/price-service/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// enable cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"POST"}
	config.MaxAge = time.Duration(12) * time.Hour
	r.Use(cors.New(config))

	v1 := r.Group("/v1")
	{
		v1.POST("/price", routes.CalculatePrice)
	}
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
