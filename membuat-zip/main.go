package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// path folder yang ingin di archive zip
	folderPath := "./files"
	// path folder dan nama file hasil archive zip
	zipFilePath := "./zipped/files.zip"

	// memanggil fungsi zipFolder
	if err := zipFolder(folderPath, zipFilePath); err != nil {
		fmt.Printf("Failed to zip folder: %v\n", err)
	}
}

// funsi zipFolder untuk membuat zip dari folder
// param 1: folderPath adalah path folder yang ingin di archive
// param 2: zipFilePath adalah path folder dan nama file hasil archive zip
// return error jika terjadi error
func zipFolder(folderPath, zipFilePath string) error {
	// membuat file zip di folder hasil archive
	fmt.Println("memulai membuat zip file...")
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	// membuat zip writer untuk menulis file zip
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// membaca isi folder yanbg ingin di archive
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return fmt.Errorf("failed to read folder: %w", err)
	}

	// looping file di folder dan menambahkan file ke zip
	for _, file := range files {
		// membuat path file yang akan di zip
		filePath := filepath.Join(folderPath, file.Name())

		// buka file yang ingin di zip
		fmt.Println("membuka file: ", filePath)
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		// Buat file baru di zip archive
		fmt.Println("menambahkan file ke zip: ", file.Name())
		writer, err := zipWriter.Create(file.Name())
		if err != nil {
			return fmt.Errorf("failed to create entry in zip file: %w", err)
		}

		// copy isi file ke zip archive
		fmt.Println("menyalin isi file ke zip: ", file.Name())
		if _, err := io.Copy(writer, file); err != nil {
			return fmt.Errorf("failed to copy file content to zip: %w", err)
		}
	}

	fmt.Println("zip file berhasil dibuat: ", zipFilePath)
	return nil
}
