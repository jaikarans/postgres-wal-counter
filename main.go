package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// Define command-line flags
	dirFlag := flag.String("d", "", "Directory path")
	sizeFlag := flag.Int64("s", 0, "Size")

	// Parse command-line arguments
	flag.Parse()

	// Validate required flags
	if *dirFlag == "" || *sizeFlag == 0 {
		fmt.Println("Usage: postgres-wal-count -d <directory_path> -s <size_in_MB>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Access the values of the flags
	dirPath := *dirFlag
	var size int64 = *sizeFlag

	// Open the specified directory
	dir, err := os.Open(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	// Read the directory entries
	fileInfos, err := dir.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	// Create a map to store file names and their size
	fileNames := make(map[string]int64)

	fmt.Printf("Start time: [%v]\n", time.Now());

	// Your logic using dirPath, size, and the opened directory
	fmt.Printf("Directory Path: %s\n", dirPath)
	fmt.Printf("Size Per Log File: %d MB\n", size)

	for true {
		for _, fileInfo := range fileInfos {
			if fileInfo.Mode().IsRegular() {
				fileName := fileInfo.Name()
				// size in MB
				fileNames[fileName] = fileInfo.Size() / (1024 * 1024)
			}	
			
		}
	
	// At 23:50:00 everyday count the file and tell print the total size in GB
	var currentTime = time.Now()
	if currentTime.Hour() == 23 && currentTime.Minute() == 50 && currentTime.Second() == 0 {
		var TotalFileNumber = int64(len(fileNames))
		fmt.Printf("Number of files %d\n", TotalFileNumber)
		fmt.Printf("[%v] Total size: %.2f\n", time.Now(), float64(float64(TotalFileNumber * size) / 1024))
		
	}

	// Loop runs every 1 minute
	time.Sleep(1000 * 60 * time.Millisecond)

	}

}
