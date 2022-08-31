package bigquery

import (
	"context"
	"os"

	"cloud.google.com/go/bigquery"
)

var client *bigquery.Client

type BigQueryClient struct {
	projectId string
	bqClient  *bigquery.Client
}

type Client interface {
	Dataset(tableId string) DataSet
	Query(query string) Query
}

func (b *BigQueryClient) Query(query string) Query {
	q := b.bqClient.Query(query)
	return &queryStruct{
		query: q,
	}

}
func (b *BigQueryClient) Dataset(tableId string) DataSet {
	ds := b.bqClient.Dataset(tableId)
	return &dataSetStruct{
		dataSet: ds,
	}
}

func SetUpToken() {
	token := os.Getenv("SERVICE_ACCOUNT_TOKEN")
	credentialsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	file, _ := os.OpenFile(credentialsFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	defer file.Close()
	_, err := file.WriteString(token)
	if err != nil {
		panic(err)
	}
}

func init() {
	SetUpToken()
	projectId := os.Getenv("PROJECT_ID")
	var err error
	client, err = bigquery.NewClient(context.Background(), projectId)

	if err != nil {
		panic(err)
	}
}

func MakeBQClient() *BigQueryClient {
	projectId := os.Getenv("PROJECT_ID")
	client, err := bigquery.NewClient(context.Background(), projectId)
	if err != nil {
		panic(err)
	}
	return &BigQueryClient{
		projectId: projectId,
		bqClient:  client,
	}
}
