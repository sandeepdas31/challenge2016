package src

import (
	"fmt"
	"net/http"
	"os"

	"github.com/distributor/constants"
	"github.com/distributor/controller"
	"github.com/distributor/models"
	"github.com/distributor/utils"
	"github.com/gin-gonic/gin"
)

type DistributorDetails struct {
	Distriburtor map[int]models.DistributionLocationInfo
}

func NewDistributorDetails() controller.DistributorService {
	return &DistributorDetails{
		Distriburtor: make(map[int]models.DistributionLocationInfo),
	}
}

func (d *DistributorDetails) AddDistributor(c *gin.Context, data models.DistributorDetailsInput, file *os.File) error {
	// check if distriutor already present
	_, found := d.Distriburtor[data.DistriburtorID]
	if found {
		c.String(http.StatusOK, "Distributor already present")
	}
	data.DistributorLocation = utils.PopulateIfEmpty(data.DistributorLocation)
	data.DistributorLocation = utils.ValidateAllLocations(c, data.DistributorLocation, file)
	d.Distriburtor[data.DistriburtorID] = data.DistributorLocation
	c.String(http.StatusOK, "Added Distributor successfully")
	return nil
}

func (d *DistributorDetails) CheckPermission(c *gin.Context, distributorPermission models.DistributorPermissions, file *os.File) error {
	if len(d.Distriburtor) == 0 {
		c.String(http.StatusOK, "No data present to check")
		return nil
	}
	val, present := d.Distriburtor[distributorPermission.DistriburtorID]
	if !present {
		c.String(http.StatusOK, "The information of the distriburtor is not present")
		return nil
	}
	distributorPermission.DistributorLocation = utils.Populate(distributorPermission.DistributorLocation)
	found := utils.ValidateLocation(c, distributorPermission.DistributorLocation, file)
	if !found {
		c.String(http.StatusBadRequest, "Invalid Location")
		return fmt.Errorf("invalid location")
	}
	fmt.Println("running exlcude")
	excludedVal := val.Exclude
	for _, distributorExclude := range excludedVal {
		if distributorPermission.DistributorLocation.Country == constants.All {
			c.String(http.StatusOK, "No")
			return nil
		}
		if distributorExclude.Country == distributorPermission.DistributorLocation.Country {
			if distributorPermission.DistributorLocation.Province == constants.All || distributorExclude.Province == constants.All {
				c.String(http.StatusOK, "No")
				return nil
			}
			if distributorExclude.Province == distributorPermission.DistributorLocation.Province {
				if distributorPermission.DistributorLocation.City == constants.All || distributorExclude.City == constants.All {
					c.String(http.StatusOK, "No")
					return nil
				}
				if distributorExclude.City == distributorPermission.DistributorLocation.City {
					c.String(http.StatusOK, "No")
					return nil
				}
			}
		}
	}
	fmt.Println("running include")
	IncludeVal := val.Include
	for _, distributorInclude := range IncludeVal {
		if distributorPermission.DistributorLocation.Country == constants.All {
			c.String(http.StatusOK, "Yes")
			return nil
		}
		if distributorInclude.Country == distributorPermission.DistributorLocation.Country {
			if distributorPermission.DistributorLocation.Province == constants.All || distributorInclude.Province == constants.All {
				c.String(http.StatusOK, "Yes")
				return nil
			}
			if distributorInclude.Province == distributorPermission.DistributorLocation.Province {
				if distributorPermission.DistributorLocation.City == constants.All || distributorInclude.City == constants.All {
					c.String(http.StatusOK, "Yes")
					return nil
				}
				if distributorInclude.City == distributorPermission.DistributorLocation.City {
					c.String(http.StatusOK, "Yes")
					return nil
				}
			}
		}
	}
	fmt.Println("running nil")
	c.String(http.StatusOK, "No")
	return nil
}

func (d *DistributorDetails) GetAllDistributors(c *gin.Context) error {
	c.IndentedJSON(http.StatusOK, d.Distriburtor)
	return nil
}

func (d *DistributorDetails) AddSubDistributor(c *gin.Context, distributorDetails models.AddSubDistributor, file *os.File) error {
	// check if the main Distributor is present
	value, found := d.Distriburtor[distributorDetails.SubDistriburtorID]
	if !found {
		c.String(http.StatusOK, "SubDistributor not found ")
		return nil
	}

	distributorDetails.Distributor.DistributorLocation = utils.PopulateIfEmpty(distributorDetails.Distributor.DistributorLocation)
	distributorDetails.Distributor.DistributorLocation = utils.ValidateAllLocations(c, distributorDetails.Distributor.DistributorLocation, file)
	var excluded bool
	var Included bool
	var updatedExcludedLocation []models.Location
	var updatedIncludedLocation []models.Location
	var distributorLocation models.Location
	updatedExcludedLocation = append(updatedExcludedLocation, value.Exclude...)
	updatedIncludedLocation = append(updatedIncludedLocation, value.Include...)

	// Check duplicates in excluded list
	for _, distributorLocation = range distributorDetails.Distributor.DistributorLocation.Exclude {
		excluded = false
		for _, location := range value.Exclude {
			if location == distributorLocation {
				fmt.Println("excluded true", location, distributorLocation)
				excluded = true
				continue
			}
		}
		fmt.Println("after continue")
		if !excluded {
			fmt.Println("running add excluded")
			updatedExcludedLocation = append(updatedExcludedLocation, distributorLocation)
		}
	}

	// Check duplicates in Included list
	for _, distributorLocation = range distributorDetails.Distributor.DistributorLocation.Include {
		Included = false
		for _, location := range value.Include {
			if location == distributorLocation {
				fmt.Println("Included true", location, distributorLocation)
				Included = true
				continue
			}
		}
		fmt.Println("after continue")
		if !Included {
			fmt.Println("running add included")
			updatedIncludedLocation = append(updatedIncludedLocation, distributorLocation)
		}
	}

	distributorDetails.Distributor.DistributorLocation.Exclude = updatedExcludedLocation
	distributorDetails.Distributor.DistributorLocation.Include = updatedIncludedLocation
	d.Distriburtor[distributorDetails.Distributor.DistriburtorID] = distributorDetails.Distributor.DistributorLocation

	c.String(http.StatusOK, "Added Distributor successfully")
	return nil
}
