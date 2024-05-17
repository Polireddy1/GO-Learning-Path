package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := ec2.NewFromConfig(cfg)

	result, err := svc.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-04ff98ccbfa41c9ad"),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	})

	if err != nil {
		log.Fatalf("Error creating instance, %v", err)
	}

	fmt.Println("Created instance", *result.Instances[0].InstanceId)
}
