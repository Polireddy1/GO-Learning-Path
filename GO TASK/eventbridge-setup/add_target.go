// // package main

// // import (
// // 	"fmt"

// // 	"github.com/aws/aws-sdk-go/aws"
// // 	"github.com/aws/aws-sdk-go/aws/session"
// // 	"github.com/aws/aws-sdk-go/service/eventbridge"
// // )

// // func AddTarget() {
// // 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// // 		SharedConfigState: session.SharedConfigEnable,
// // 	}))

// // 	svc := eventbridge.New(sess)

// // 	target := &eventbridge.Target{
// // 		Arn: aws.String("arn:aws:lambda:ap-south-1-2:123456789012:function:my-function"),
// // 		Id:  aws.String("MyLambdaFunction"),
// // 	}

// // 	input := &eventbridge.PutTargetsInput{
// // 		Rule:         aws.String("ECRPushEventRule"),
// // 		Targets:      []*eventbridge.Target{target},
// // 		EventBusName: aws.String("my-event-bus"),
// // 	}

// // 	_, err := svc.PutTargets(input)
// // 	if err != nil {
// // 		fmt.Println("Error adding target to rule:", err)
// // 		return
// // 	}

// // 	fmt.Println("Target added to rule successfully")
// // }


// package main

// import (
//     "fmt"
//     "github.com/aws/aws-sdk-go/aws"
//     "github.com/aws/aws-sdk-go/aws/session"
//     "github.com/aws/aws-sdk-go/service/eventbridge"
// )

// func AddTarget() {
//     sess := session.Must(session.NewSession(&aws.Config{
//         Region: aws.String("ap-south-1"),
//     }))

//     svc := eventbridge.New(sess)

//     target := &eventbridge.Target{
//         Arn: aws.String("arn:aws:lambda:ap-south-1:851725240994:function:lambda-function"),
//         Id:  aws.String("MyLambdaFunction"),
//     }

//     input := &eventbridge.PutTargetsInput{
//         Rule:        aws.String("ECRPushEventRule"),
//         Targets:     []*eventbridge.Target{target},
//         EventBusName: aws.String("my-event-bus"),
//     }

//     _, err := svc.PutTargets(input)
//     if err != nil {
//         fmt.Println("Error adding target to rule:", err)
//         return
//     }

//     fmt.Println("Target added to rule successfully")
// }


