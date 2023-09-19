package handler

import (
	"fmt"
	"os"
)

func ReadCSV() (*os.File, error) {
	fmt.Println("read csv file ")
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return nil, err
	}
	fmt.Println("currentDir", currentDir)
	filePath := currentDir + "/cities.csv"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return file, nil
}
