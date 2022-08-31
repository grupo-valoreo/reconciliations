package bigquery

import (
	"cloud.google.com/go/bigquery"
)

type Table interface {
	Inserter() Inserter
}

type tableStruct struct {
	table *bigquery.Table
}

func (o *tableStruct) Inserter() Inserter {
	inserter := o.table.Inserter()
	return &inserterStruct{
		inserter: inserter,
	}
}
