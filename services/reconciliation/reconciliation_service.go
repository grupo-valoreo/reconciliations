package reconciliation_service

import (
	"context"

	bigqueryDomain "github.com/grupo-valoreo/reconciliations/domain/bigquery"

	bq "github.com/grupo-valoreo/reconciliations/repositories/bigquery"
)

type ReconciliationService interface {
	Get_Aju_Trans(ctx context.Context, q string) string
	Store_Aju_Trans(ctx context.Context, data interface{}, dataset *bigqueryDomain.Dataset) error
}

type reconciliationService struct {
	bqRepo bq.BigQueryRepository
}

func MakeBigQueryService(bqRepo bq.BigQueryRepository) ReconciliationService {

	return &reconciliationService{
		bqRepo: bqRepo,
	}

}

func (s *reconciliationService) Store_Aju_Trans(ctx context.Context, data interface{}, dataset *bigqueryDomain.Dataset) error {
	s.bqRepo.Store(data, dataset)
	return s.bqRepo.Store(data, dataset)
}

func (s *reconciliationService) Get_Aju_Trans(ctx context.Context, q string) string {
	s.bqRepo.Get(ctx, q)
	return "golas"
}
