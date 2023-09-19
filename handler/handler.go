package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/distributor/controller"
	"github.com/distributor/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	handler controller.DistributorService
	log     *logrus.Logger
	file    *os.File
}

func NewHandler(h controller.DistributorService, Log *logrus.Logger, File *os.File) *Handler {
	return &Handler{
		handler: h,
		log:     Log,
		file:    File,
	}
}

func (h *Handler) GetAllDistributors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.handler.GetAllDistributors(c))
}

func (h *Handler) AddDistributor(c *gin.Context) {
	var newDistributor models.DistributorDetailsInput
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	err = json.Unmarshal(body, &newDistributor)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	log.Println("distributors", newDistributor)
	err = h.handler.AddDistributor(c, newDistributor, h.file)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.String(200, fmt.Sprintf("%#v", newDistributor))
}

func (h *Handler) CheckPermission(c *gin.Context) {
	log.Println("checking permissions")
	var newDistributor models.DistributorPermissions
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &newDistributor)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	err = h.handler.CheckPermission(c, newDistributor, h.file)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

}

func (h *Handler) AddSubDistributor(c *gin.Context) {
	var newDistributor models.AddSubDistributor
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	err = json.Unmarshal(body, &newDistributor)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	log.Println("distributors", newDistributor)
	err = h.handler.AddSubDistributor(c, newDistributor, h.file)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.String(200, fmt.Sprintf("%#v", newDistributor))
}
