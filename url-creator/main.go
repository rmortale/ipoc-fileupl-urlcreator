package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (string, error) {
	return "Hello Go", nil
}

func main() {
	lambda.Start(handler)
}
