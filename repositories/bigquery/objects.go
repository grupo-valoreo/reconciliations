package bigquery

import (
	"context"

	"github.com/grupo-valoreo/reconciliations/domain/bigquery"
)

func (repo *bigQueryRepository) Store(data interface{}, dataset *bigquery.Dataset) error {
	dataSet := repo.client.Dataset(dataset.DataSet)
	table := dataSet.Table(dataset.Name)

	inserter := table.Inserter()

	//inserter.IgnoreUnknownValues = true
	return inserter.Put(context.Background(), data)
}

func (repo *bigQueryRepository) Get(ctx context.Context, q string) interface{} {
	query := repo.client.Query(q)
	data := query.Read(ctx)
	//inserter.IgnoreUnknownValues = true
	return data
}
