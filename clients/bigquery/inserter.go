package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type Inserter interface {
	Put(ctx context.Context, src interface{}) error
}

type inserterStruct struct {
	inserter *bigquery.Inserter
}

func (w *inserterStruct) Put(ctx context.Context, src interface{}) error {
	return w.inserter.Put(ctx, src)
}
