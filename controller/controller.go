package controller

import (
	"os"

	"github.com/distributor/models"
	"github.com/gin-gonic/gin"
)

type DistributorService interface {
	AddDistributor(*gin.Context, models.DistributorDetailsInput, *os.File) error
	CheckPermission(*gin.Context, models.DistributorPermissions, *os.File) error
	GetAllDistributors(*gin.Context) error
	AddSubDistributor(*gin.Context, models.AddSubDistributor, *os.File) error
}
