package main

import (
    "context"
    "fmt"
    "log"
    //"github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }

    svc := ec2.NewFromConfig(cfg)

    instanceId := "i-013363584bd3b3513" 
    input := &ec2.StopInstancesInput{
        InstanceIds: []string{instanceId},
    }

    result, err := svc.StopInstances(context.TODO(), input)
    if err != nil {
        log.Fatalf("Error stopping instance: %v", err)
    }

    fmt.Println("Instance stopped successfully:", result.StoppingInstances)
}
