// go get github.com/awslabs/aws-sdk-go/service/dynamodb
package main

import (
	"fmt"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)

func main() {
	svc := dynamodb.New(&aws.Config{Region: "ap-southeast-1"})

	params := &dynamodb.GetItemInput{
		Key: &map[string]*dynamodb.AttributeValue{ // Required
			"Email": &dynamodb.AttributeValue{ // Required
				S: aws.String("warren.gutierrez@gmail.com"),
			},
		},
		TableName:      aws.String("Users"),
		ConsistentRead: aws.Boolean(true),
	}
	resp, err := svc.GetItem(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}
