package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/distributor/handler"
	"github.com/distributor/src"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("running rest api")

	//handler.Init()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	var log *logrus.Logger
	DistributorHandler := src.NewDistributorDetails()
	file, err := handler.ReadCSV()
	if err != nil {
		log.Error(err)
	}
	handler := handler.NewHandler(DistributorHandler, log, file)
	r.GET("/all-distributors", handler.GetAllDistributors)
	r.POST("/add-distributors", handler.AddDistributor)
	r.POST("/add-sub-distributors", handler.AddSubDistributor)
	r.GET("/check-permission", handler.CheckPermission)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	defer file.Close()
}
