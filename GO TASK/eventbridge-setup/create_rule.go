// // package main

// // import (
// //     "fmt"
// //     "github.com/aws/aws-sdk-go/aws"
// //     "github.com/aws/aws-sdk-go/aws/session"
// //     "github.com/aws/aws-sdk-go/service/eventbridge"
// // )

// // func CreateRule() {
// //     sess := session.Must(session.NewSessionWithOptions(session.Options{
// //         SharedConfigState: session.SharedConfigEnable,
// //     }))

// //     svc := eventbridge.New(sess)

// //     eventPattern := `{
// //         "source": ["aws.ecr"],
// //         "detail-type": ["ECR Image Action"]
// //     }`

// //     input := &eventbridge.PutRuleInput{
// //         Name:         aws.String("ECRPushEventRule"),
// //         EventPattern: aws.String(eventPattern),
// //         EventBusName: aws.String("my-event-bus"),
// //     }

// //     result, err := svc.PutRule(input)
// //     if err != nil {
// //         fmt.Println("Error creating rule:", err)
// //         return
// //     }

// //     fmt.Println("Rule ARN:", *result.RuleArn)
// // }


// package main

// import (
//     "fmt"
//     "github.com/aws/aws-sdk-go/aws"
//     "github.com/aws/aws-sdk-go/aws/session"
//     "github.com/aws/aws-sdk-go/service/eventbridge"
// )

// func CreateRule() {
//     sess := session.Must(session.NewSession(&aws.Config{
//         Region: aws.String("ap-south-1"),
//     }))

//     svc := eventbridge.New(sess)

//     eventPattern := `{
//         "source": ["aws.ecr"],
//         "detail-type": ["ECR Image Action"]
//     }`

//     input := &eventbridge.PutRuleInput{
//         Name:         aws.String("ECRPushEventRule"),
//         EventPattern: aws.String(eventPattern),
//         EventBusName: aws.String("my-event-bus"),
//     }

//     result, err := svc.PutRule(input)
//     if err != nil {
//         fmt.Println("Error creating rule:", err)
//         return
//     }

//     fmt.Println("Rule ARN:", *result.RuleArn)
// }
