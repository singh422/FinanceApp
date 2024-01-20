package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Reader struct {
	directoryPath string
}

func (r *Reader) readCSVFiles(directoryPath string) error {
	files, err := r.getCSVFiles(directoryPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Printf("Reading file: %s\n", file)
		err := r.readCSVFile(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Reader) getCSVFiles(directoryPath string) ([]string, error) {
	var csvFiles []string

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".csv" {
			csvFiles = append(csvFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return csvFiles, nil
}

func (r *Reader) readCSVFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		fmt.Println(record)
	}

	return nil
}

func main() {
	directoryPath := "/path/to/csv/files"
	err := readCSVFiles(directoryPath)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
