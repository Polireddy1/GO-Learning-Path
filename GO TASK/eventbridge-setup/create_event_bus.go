// // package main

// // import (
// //     "fmt"
// //     "github.com/aws/aws-sdk-go/aws"
// //     "github.com/aws/aws-sdk-go/aws/session"
// //     "github.com/aws/aws-sdk-go/service/eventbridge"
// // )

// // func CreateEventBus() {
// //     sess := session.Must(session.NewSessionWithOptions(session.Options{
// //         SharedConfigState: session.SharedConfigEnable,
// //     }))

// //     svc := eventbridge.New(sess)

// //     input := &eventbridge.CreateEventBusInput{
// //         Name: aws.String("my-event-bus"),
// //     }

// //     result, err := svc.CreateEventBus(input)
// //     if err != nil {
// //         fmt.Println("Error creating event bus:", err)
// //         return
// //     }

// //     fmt.Println("Event bus ARN:", *result.EventBusArn)
// // }


// package main

// import (
//     "fmt"
//     "github.com/aws/aws-sdk-go/aws"
//     "github.com/aws/aws-sdk-go/aws/session"
//     "github.com/aws/aws-sdk-go/service/eventbridge"
// )

// func CreateEventBus() {
//     sess := session.Must(session.NewSession(&aws.Config{
//         Region: aws.String("ap-south-1"),
//     }))

//     svc := eventbridge.New(sess)

//     input := &eventbridge.CreateEventBusInput{
//         Name: aws.String("my-event-bus"),
//     }

//     result, err := svc.CreateEventBus(input)
//     if err != nil {
//         fmt.Println("Error creating event bus:", err)
//         return
//     }

//     fmt.Println("Event bus ARN:", *result.EventBusArn)
// }

