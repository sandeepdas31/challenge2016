package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/distributor/constants"
	"github.com/distributor/models"
	"github.com/gin-gonic/gin"
)

//var wg sync.WaitGroup
/* to check correct location using routines */
// func ValidateAllLocations(location models.DistributionLocationInfo, reader *csv.Reader) models.DistributionLocationInfo {
// 	fmt.Println("Running Validate")
// 	var updatedExcludedLocation []models.Location
// 	var updatedIncludedLocation []models.Location
// 	excludeChannel := make(chan map[bool]models.Location)
// 	for _, value := range location.Exclude {
// 		wg.Add(1)
// 		go ValidateLocation(value, reader, excludeChannel, &wg)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(excludeChannel)
// 	}()
// 	for exclude := range excludeChannel {
// 		fmt.Println("exclude", exclude)
// 		for key, value := range exclude {
// 			if key {
// 				updatedExcludedLocation = append(updatedExcludedLocation, value)
// 			}
// 		}
// 	}
// 	// for _, value := range location.Include {
// 	// 	found := ValidateLocation(&value, reader)
// 	// 	if !found {
// 	// 		fmt.Println("Invalid include location", value)
// 	// 	} else {
// 	// 		updatedIncludedLocation = append(updatedIncludedLocation, value)
// 	// 	}
// 	// }

// 	includeChannel := make(chan map[bool]models.Location)
// 	for _, value := range location.Include {
// 		wg.Add(1)
// 		go ValidateLocation(value, reader, includeChannel, &wg)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(includeChannel)
// 	}()
// 	for include := range includeChannel {
// 		fmt.Println("include", include)
// 		for key, value := range include {
// 			if key {
// 				updatedIncludedLocation = append(updatedIncludedLocation, value)
// 			}
// 		}
// 	}

// 	fmt.Println("updatedIncludedLocation", updatedIncludedLocation, "updatedExcludedLocation", updatedExcludedLocation)
// 	return models.DistributionLocationInfo{
// 		Include: updatedIncludedLocation,
// 		Exclude: updatedExcludedLocation,
// 	}
// }

func ValidateAllLocations(c *gin.Context, location models.DistributionLocationInfo, file *os.File) models.DistributionLocationInfo {
	fmt.Println("Running Validate")
	var updatedExcludedLocation []models.Location
	var updatedIncludedLocation []models.Location

	for _, value := range location.Exclude {
		found := ValidateLocation(c, value, file)
		if found {
			updatedExcludedLocation = append(updatedExcludedLocation, value)
		} else {
			c.String(http.StatusOK, "Invalid location")
			c.IndentedJSON(http.StatusOK, value)
		}
	}
	for _, value := range location.Include {
		found := ValidateLocation(c, value, file)
		if found {
			updatedIncludedLocation = append(updatedIncludedLocation, value)
		} else {
			c.String(http.StatusOK, "Invalid location")
			c.IndentedJSON(http.StatusOK, value)
		}
	}
	return models.DistributionLocationInfo{
		Exclude: updatedExcludedLocation,
		Include: updatedIncludedLocation,
	}
}

func ValidateLocation(c *gin.Context, location models.Location, file *os.File) bool {
	fmt.Println("running validate location", location)
	//defer wg.Done()
	var found = false
	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Printf("Error seeking to the beginning of the file: %v\n", err)
		return false
	}
	reader := csv.NewReader(file)
	for {
		fmt.Println("Inside reader")
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				fmt.Println("end of file")
				break
			} else {
				fmt.Printf("Error reading CSV: %v\n", err)
				return false
			}
		}

		// Process the current line
		fmt.Println("strings.ToLower(record[3])", strings.ToLower(record[3]), strings.ToLower(record[4]), strings.ToLower(record[5]))
		// Check if the data is present in the desired columns
		if (strings.ToLower(record[3]) == location.City || location.City == constants.All) &&
			(strings.ToLower(record[4]) == location.Province || location.Province == constants.All) &&
			(strings.ToLower(record[5]) == location.Country || location.Country == constants.All) {
			fmt.Println("found", strings.ToLower(record[3]), strings.ToLower(record[4]), strings.ToLower(record[5]))
			found = true
			break
		}

		// If you want to process one line at a time, you can break her
	}

	//Channel <- map[bool]models.Location{found: location}
	return found
}

func PopulateIfEmpty(location models.DistributionLocationInfo) models.DistributionLocationInfo {
	fmt.Println("PopulateIfEmpty")
	for key, value := range location.Include {
		value = Populate(value)
		location.Include[key] = value
	}

	for key, value := range location.Exclude {
		value = Populate(value)
		location.Exclude[key] = value
	}
	return location
}

func Populate(location models.Location) models.Location {
	if location.Country == "" {
		fmt.Println("Please populate country")
	}
	if location.Province == "" {
		location.Province = constants.All
	}
	if location.City == "" {
		location.City = constants.All
	}

	return ToLowerAndTrim(location)
}

func ToLowerAndTrim(location models.Location) models.Location {
	regex := regexp.MustCompile(`\s+`)
	location.Country = regex.ReplaceAllString(strings.ToLower(location.Country), " ")
	location.Province = regex.ReplaceAllString(strings.ToLower(location.Province), " ")
	location.City = regex.ReplaceAllString(strings.ToLower(location.City), " ")
	return location
}
