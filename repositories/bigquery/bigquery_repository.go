package bigquery

import (
	"context"

	"github.com/grupo-valoreo/reconciliations/clients/bigquery"
	bigqueryDomain "github.com/grupo-valoreo/reconciliations/domain/bigquery"
)

type BigQueryRepository interface {
	Objects
}

type Objects interface {
	Store(data interface{}, dataset *bigqueryDomain.Dataset) error
	Get(ctx context.Context, q string) interface{}
}

type bigQueryRepository struct {
	client bigquery.BigQueryClient
}

func MakeBigQueryRepository(client bigquery.BigQueryClient) BigQueryRepository {
	return &bigQueryRepository{
		client: client,
	}
}
