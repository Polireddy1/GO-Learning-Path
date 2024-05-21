// package main

// import (
// 	"archive/zip"
// 	"io"
// 	"os"
// 	"path/filepath"
// )

// func createZip(sourceFiles []string, output string) error {
// 	zipFile, err := os.Create(output)
// 	if err != nil {
// 		return err
// 	}
// 	defer zipFile.Close()

// 	zipWriter := zip.NewWriter(zipFile)
// 	defer zipWriter.Close()

// 	for _, file := range sourceFiles {
// 		err = addFileToZip(zipWriter, file)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func addFileToZip(zipWriter *zip.Writer, filename string) error {
// 	fileToZip, err := os.Open(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer fileToZip.Close()

// 	// Get the file information
// 	info, err := fileToZip.Stat()
// 	if err != nil {
// 		return err
// 	}

// 	// Create a header based on the file info
// 	header, err := zip.FileInfoHeader(info)
// 	if err != nil {
// 		return err
// 	}

// 	// Ensure the header is set to the correct path in the archive
// 	header.Name = filepath.Base(filename)
// 	header.Method = zip.Deflate

// 	// Create a writer for the file header
// 	writer, err := zipWriter.CreateHeader(header)
// 	if err != nil {
// 		return err
// 	}

// 	// Copy the file contents to the writer
// 	_, err = io.Copy(writer, fileToZip)
// 	return err
// }
