/*

This package contains code to take in a playlist id and
return playlist metadata

*/
package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nimitpatel26/yt-data-fetcher/go/ytData"
)

func getPlaylistData(playlistId string) string {
	data, _ := json.Marshal(ytData.GetPlayListData(playlistId))
	return string(data)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       getPlaylistData(request.QueryStringParameters["id"]),
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
