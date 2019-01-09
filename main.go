package main

import (
	"encoding/json"

	rz "github.com/adrianpk/rezerw/core"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ct    = "Content-Type"
	ctVal = "application/json"
)

func findOne(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := r.PathParameters["id"]
	// Find
	account, err := rz.FindAccountByIDString(id)
	if err != nil {
		return rz.RenderError(err)
	}
	// Marshall
	res, err := json.Marshal(account)
	if err != nil {
		return rz.RenderError(err)
	}
	// Response
	return rz.RenderOk(res)

}

func main() {
	lambda.Start(findOne)
}
