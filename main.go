package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/deibyssoca/awsgo"
)

func main() {
	lambda.Start(executeLambda)
}

func executeLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializeAWS()

	return event, ctx.Err()
}

func ValidateParam() bool {
	return true
}
