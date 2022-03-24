package main

import "github.com/aws/aws-lambda-go/lambda"

func displayThis() (string, error) {
	return "Serverless Software Development using AWS Lambda with Go", nil
}

func main() {
	lambda.Start(displayThis)
}
