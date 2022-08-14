package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"lambda-store-nir/service/application/service"
	"lambda-store-nir/service/infraestructure/s3"
	"net/http"
)

func hfn(event events.S3Event) {

}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	println(req.Path)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "text/plain; charset=utf-8"},
		Body:       "Olá mundo",
	}, nil

}

func main() {

	store := s3.NewStoreS3()
	service.NewDocumentService(store)

	//calc(2, 6.9)
	lambda.Start(handler)
}
