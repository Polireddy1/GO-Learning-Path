// // zippp.go
// package main

// import (
// 	"log"

// 	"./create_zip"
// )
// func main() {
// 	sourceFiles := []string{"bootstrap", "main"}
// 	zipFilePath := "function.zip"
// 	// functionName := "main" // Replace with your actual Lambda function name

// 	err := create_zip.CreateZip(sourceFiles, zipFilePath)
// 	if err != nil {
// 		log.Fatalf("Failed to create ZIP file: %v", err)
// 	}

// 	// err = upload_lambda.UploadFunction(zipFilePath, functionName)
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to upload Lambda function: %v", err)
// 	// }
// }
