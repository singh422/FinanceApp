package reader

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"

	"github.com/singh422/FinanceApp/internal/source"
	"github.com/singh422/FinanceApp/internal/util"
)

type Reader struct {
	DirectoryPath string
}

type FileInfo struct {
	Records  [][]string
	Source   source.Source
	FileName string
}

func (r *Reader) ReadCSVFiles(directoryPath string) ([]*FileInfo, error) {

	var fileInfos []*FileInfo
	files, err := r.getCSVFiles(directoryPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		// fmt.Printf("Reading file: %s\n", file)
		fileInfo, err := r.readCSVFile(file)
		if err != nil {
			return nil, err
		}

		if fileInfo.Source == source.Unknown {
			log.Printf("WARNING:Skipping file: [%s], cannot identify source", file)
		} else {
			fileInfos = append(fileInfos, fileInfo)
			log.Printf("INFO:Successfully read file [%s] for source [%s]. Found [%d] records.", file, fileInfo.Source, len(fileInfo.Records))
		}
	}

	return fileInfos, nil
}

func (r *Reader) getCSVFiles(directoryPath string) ([]string, error) {
	var csvFiles []string

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info != nil && info.IsDir() && info.Name() == "Output Reports" {
			return filepath.SkipDir
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

func (r *Reader) readCSVFile(filePath string) (*FileInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return &FileInfo{
		FileName: filePath,
		Records:  records,
		Source:   parseSourceFromFileName(filePath),
	}, nil
}

func parseSourceFromFileName(filePath string) source.Source {
	sources := source.GetAllSources()

	for _, src := range sources {
		if util.CaseInsensitiveSubstring(filePath, src.String()) {
			return src
		}
	}

	return source.Unknown
}
