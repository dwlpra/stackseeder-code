package main

import (
	"archive/zip"
	"io"
	"os"
	"path"
)

func main() {

	folderPath := "./files"
	zipFilePath := path.Join("./zipped/files.zip")

	err := zipFolder(folderPath, zipFilePath)
	if err != nil {
		panic(err)
	}
}

func zipFolder(folderPath, zipFilePath string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := path.Join(folderPath, file.Name())
		if err != nil {
			return err
		}
		f, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer f.Close()

		w, err := zipWriter.Create(file.Name())
		if err != nil {
			return err
		}
		if _, err := io.Copy(w, f); err != nil {
			return err
		}
	}
	return nil
}
