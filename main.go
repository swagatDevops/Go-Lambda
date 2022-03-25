package main

import "github.com/aws/aws-lambda-go/lambda"

func Handler() (string, error) {
	return "Serverless Software Development using AWS Lambda with Go", nil
}

func main() {
	lambda.Start(Handler)
}
