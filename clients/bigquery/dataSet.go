package bigquery

import (
	"cloud.google.com/go/bigquery"
)

type DataSet interface {
	Table(tableID string) Table
}

type dataSetStruct struct {
	dataSet *bigquery.Dataset
}

func (d *dataSetStruct) Table(tableID string) Table {
	table := d.dataSet.Table(tableID)
	return &tableStruct{
		table: table,
	}
}
