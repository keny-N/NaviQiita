package main

import (
	"context"

	"githun.com/keny-N/NaviQiita/internal/qiita"
	"githun.com/keny-N/NaviQiita/pkg/response"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context) (response.AlexaResponse, error) {
	items, err := qiita.FetchQiitaItems()
	if err != nil {
		return response.AlexaResponse{}, err
	}

	message := "最新のQiitaの記事を５つ紹介します。。: "
	for _, item := range items {
		message += item.Title + ", URL: " + item.URL + ". "
	}

	return response.CreateResponse(message), nil
}

func main() {
	lambda.Start(Handler)
}
