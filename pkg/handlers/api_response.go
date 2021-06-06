package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)


    
func apiResponse(status int, body interface{}) (events.APIGatewayProxyResponse, error) {

  stringBody, _ := json.Marshal(body)

  resp := events.APIGatewayProxyResponse{
		StatusCode:      status,
		IsBase64Encoded: false,
		Body:            string(stringBody),
		Headers: map[string]string{
			"Content-Type":           "application/json",
		},
	}


  return resp, nil
}