package app

import (
	"github.com/grupo-valoreo/reconciliations/clients/bigquery"
	reconciliation_impl "github.com/grupo-valoreo/reconciliations/implementation/reconciliation"
	bqRepo "github.com/grupo-valoreo/reconciliations/repositories/bigquery"
	bqService "github.com/grupo-valoreo/reconciliations/services/reconciliation"
)

func (app *App) includeRoutes() {

	//Clients
	bigQueryClient := bigquery.MakeBQClient()
	//Repo
	bigQueryRepo := bqRepo.MakeBigQueryRepository(*bigQueryClient)
	//Servicio
	reconciliationService := bqService.MakeBigQueryService(bigQueryRepo)
	//Implementations
	reconciliationImpl := reconciliation_impl.MakeReconciliationImpl(reconciliationService)
	//Add APIs
	app.AddApi(reconciliationImpl)
}
