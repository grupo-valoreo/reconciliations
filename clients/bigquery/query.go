package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type Query interface {
	Read(ctx context.Context) bigquery.RowIterator
}

type queryStruct struct {
	query *bigquery.Query
}

func (q *queryStruct) Read(ctx context.Context) bigquery.RowIterator {
	row, error := q.query.Read(context.Background())
	println(error)
	return *row
}
