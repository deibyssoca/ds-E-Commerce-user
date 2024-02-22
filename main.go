package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/deibyssoca/awsgo"
	"github.com/deibyssoca/bd"
	"github.com/deibyssoca/models"
)

func main() {
	lambda.Start(executeLambda)
}

func executeLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializeAWS()
	if !ValidateParam() {
		fmt.Println("Error in the parameters. The 'SecretName' must exist")
		err := errors.New("error in the parameters. the secret name must exist")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Emails = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	if err := bd.ReadSecret(); err != nil {
		fmt.Println("Error when reading the secret." + err.Error())
		return event, err
	}

	return event, bd.SignUp(data)
}

// valido que me manden el parametro
func ValidateParam() bool {
	var parameterExist bool
	_, parameterExist = os.LookupEnv("SecretName")
	return parameterExist
}
