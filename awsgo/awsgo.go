package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var ctx context.Context
var conf aws.Config
var err error

func InicializeAWS() {
	/* Devuelve un contexto vacío no nulo  */
	ctx = context.TODO()
	/* Usamos esta región porque no es necesario cambiar, si fuera necesario podríamos tener la region en una variable de entorno */
	conf, err = config.LoadDefaultConfig(ctx, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		panic("Error al cargar la configurations .aws/config " + err.Error())
	}
}
