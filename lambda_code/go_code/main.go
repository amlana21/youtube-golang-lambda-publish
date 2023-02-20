package main

import (
	"context"
	"fmt"

	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type LambdaEvent struct {
	Type string `json:"type"`
}

// Create struct to hold info about User
type User struct {
	Username string
	Name     string
	Status   string
}

type LambdaResponse struct {
	Message string `json:"message"`
}

func LambdaHandler(ctx context.Context, event LambdaEvent) ([]User, error) {

	// Initiate session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Declare Dynamodb svc
	svc := dynamodb.New(sess)
	tableName := "Users"
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	// Query DB
	result, err := svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}
	responseArray := []User{}

	for _, i := range result.Items {
		item := User{}

		err = dynamodbattribute.UnmarshalMap(i, &item)
		responseArray = append(responseArray, item)

		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}
		fmt.Println("Username: ", item.Username)
		fmt.Println("Name:", item.Name)
		fmt.Println()
	}
	return responseArray, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
