package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	uploader := manager.NewUploader(s3.NewFromConfig(cfg))

	f, err := os.Open("docker.txt") 
	if err != nil {
		log.Fatalf("failed to open file %q, %v", "file.txt", err)
	}
	defer f.Close()

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("lepsbucket"), 
		Key:    aws.String("docker.txt"),
		Body:   f,
	})
	if err != nil {
		log.Fatalf("failed to upload file, %v", err)
	}

	fmt.Printf("file uploaded to, %s\n", result.Location)
}
