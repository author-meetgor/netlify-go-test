package main

import (
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    return &events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body: "Hello, Gophers!",
    }, nil
}

func main() {
    lambda.Start(Handler)
}
