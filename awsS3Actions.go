package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	creds := aws.Credentials{
		AccessKeyID:     "",
		SecretAccessKey: "",
	}

	// Create a static credentials provider
	credsProvider := credentials.StaticCredentialsProvider{
		Value: creds,
	}

	// Load AWS SDK configuration with static credentials and region
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credsProvider),
		config.WithRegion("us-east-1")) // Replace "us-east-1" with your desired AWS region

	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("testing-task1"),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bucket Contents:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}

	file, err := os.Open("C:/Users/Relanto/Desktop/file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()

	// Upload the file to S3
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String("testing-task1"),
		Key:           aws.String("file2.txt"), // Replace with the desired file path in the bucket
		Body:          file,
		ContentLength: aws.Int64(fileSize),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File uploaded successfully!")

	// Check if there are any objects in the bucket
	if len(output.Contents) > 0 {
		// Get the key (file path) of the first object
		firstObjectKey := aws.ToString(output.Contents[0].Key)

		// Download the first object
		obj, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String("testing-task1"),
			Key:    aws.String(firstObjectKey),
		})
		if err != nil {
			log.Fatal(err)
		}
		defer obj.Body.Close()

		// Read the object's contents
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, obj.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Print the contents of the first file
		fmt.Printf("Contents of the first file (%s):\n%s\n", firstObjectKey, buf.String())
	} else {
		fmt.Println("The bucket is empty.")
	}
}
