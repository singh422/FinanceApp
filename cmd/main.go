package main

import "fmt"

func main() {
	directoryPath := "/path/to/csv/files"
	err := readCSVFiles(directoryPath)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
