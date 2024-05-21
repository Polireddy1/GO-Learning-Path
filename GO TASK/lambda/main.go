package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type ECRDetail struct {
	ImageTag string `json:"image-tag"`
}

type ECREvent struct {
	Detail ECRDetail `json:"detail"`
}

func ExtractImageTag(event []byte) (string, error) {
	var ecrEvent ECREvent
	if err := json.Unmarshal(event, &ecrEvent); err != nil {
		return "", fmt.Errorf("failed to unmarshal event: %v", err)
	}
	return ecrEvent.Detail.ImageTag, nil
}

func handleRequest(ctx context.Context, event []byte) (string, error) {
	imageTag, err := ExtractImageTag(event)
	if err != nil {
		return "", err
	}
	if imageTag == "" {
		return "", fmt.Errorf("image tag is missing in the event")
	}

	return fmt.Sprintf("Extracted image tag: %s", imageTag), nil
}

func main() {
	lambda.Start(handleRequest)
}
