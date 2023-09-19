package constants

var (
	Country  = 3
	Province = 4
	City     = 5
	All      = "all"
)

// var found bool

// 	// Start processing lines in parallel
// 	const numWorkers = 4 // Adjust the number of workers as needed
// 	lines := make(chan []string, numWorkers)

// 	// Start workers
// 	for i := 0; i < numWorkers; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for line := range lines {
// 				// Check if the data is present in the desired columns
// 				if len(line) >= 3 && (line[0] == searchData || line[1] == searchData || line[2] == searchData) {
// 					found = true
// 					close(lines) // Close the channel to stop other workers
// 					return
// 				}
// 			}
// 		}()
// 	}

// 	// Read the CSV records and distribute them to workers
// 	for {
// 		record, err := reader.Read()
// 		if err != nil {
// 			// End of file
// 			if err == csv.ErrEOF {
// 				break
// 			}
// 			fmt.Printf("Error reading CSV: %v\n", err)
// 			return
// 		}
// 		lines <- record
// 	}

// 	// Wait for all workers to finish
// 	close(lines)
// 	wg.Wait()

// 	// Check if data was found
// 	if found {
// 		fmt.Printf("Data '%s' is present in the CSV.\n", searchData)
// 	} else {
// 		fmt.Printf("Data '%s' is not found in the CSV.\n", searchData)
// 	}
// }
