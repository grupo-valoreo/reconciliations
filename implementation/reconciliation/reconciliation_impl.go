package reconciliation_impl

import (
	"context"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	reconciliation_service "github.com/grupo-valoreo/reconciliations/services/reconciliation"
)

type ReconciliationImpl interface {
	Reconciliate(ctx context.Context) struct{}
}

type reconciliationImpl struct {
	reconciliationService reconciliation_service.ReconciliationService
}

func (*reconciliationImpl) Routes() map[string]string {
	return map[string]string{
		"Reconciliate": "/reconciliate",
	}
}

func MakeReconciliationImpl(reconciliation reconciliation_service.ReconciliationService) ReconciliationImpl {
	return &reconciliationImpl{
		reconciliationService: reconciliation,
	}
}

func (r *reconciliationImpl) Reconciliate(ctx context.Context) struct{} {

	// dataset := bigqueryDomain.Dataset{
	// 	DataSet: "testTerraformDS",
	// 	Name:    "test",
	// }

	// var datita struct {
	// 	name string
	// }
	//To bigq
	// datita.name = "tst2"
	// query := "SELECT *  FROM `valoreo-production.testTerraformDS.test` LIMIT 10"

	// a := r.reconciliationService.Store_Aju_Trans(ctx, datita, &dataset)
	// b := r.reconciliationService.Get_Aju_Trans(ctx, query)
	// fmt.Println(a)
	// fmt.Println(b)
	dataEntSal := `
		[
		{
		  "DELIVERY": 359625,
		  "DOCUMENTO": "TO98",
		  "TIPO_DOC": "itemfulfillment",
		  "IDENTIFICADOR_MOV": "TO98",
		  "FECHA_CREACION": "2022-04-19",
		  "MONEDA": null,
		  "LINEA": 11,
		  "CANTIDAD": 16,
		  "PRECIO": null,
		  "NO_PEDIMENTO": "",
		  "FECHA_PEDIMENTO": "",
		  "NOMBRE_ADUANA": "",
		  "RETORNO_BODEGA": "",
		  "ESTADO_ALMACEN": "pick",
		  "FECHA_MODIFICACION": "2022-04-19",
		  "STATUS_RESPUESTA": 200,
		  "Id busqueda": "TO98-7708113176140"
		},
		{
		  "DELIVERY": 112340,
		  "DOCUMENTO": "PO103",
		  "TIPO_DOC": "itemreceipt",
		  "IDENTIFICADOR_MOV": "6075163",
		  "FECHA_CREACION": "2022-04-28",
		  "MONEDA": 1,
		  "LINEA": 4,
		  "CANTIDAD": 48,
		  "PRECIO": "479.64",
		  "NO_PEDIMENTO": "",
		  "FECHA_PEDIMENTO": "",
		  "NOMBRE_ADUANA": "470-AICM",
		  "RETORNO_BODEGA": "returned",
		  "ESTADO_ALMACEN": "",
		  "FECHA_MODIFICACION": "2022-04-28",
		  "STATUS_RESPUESTA": 200,
		  "Id busqueda": "6075163-7500764000108"
		}
	   ]
	`

	dataLine2 := `[
		{
		  "Id": 359625,
		  "Formula": "359625-11",
		  "Item": 7708113176140
		},
		{
		  "Id": 112340,
		  "Formula": "112340-4",
		  "Item": 7500764000108
		}
	   ]`

	dataMov := `[
		{
		  "record_type": "inventoryadjustment",
		  "id": 1139363,
		  "item_internal_id": 7708113176140,
		  "custbody_val_transaction_id": "TO98",
		  "internal_id": 2490,
		  "formula_numeric": 16,
		  "formula_text": "A",
		  "subsidiary_id": null,
		  "subsidiary_name": null,
		  "type": null,
		  "creation_date": "2022-08-29"
		},
		{
		  "record_type": "inventoryadjustment",
		  "id": 756583,
		  "item_internal_id": 7500764000108,
		  "custbody_val_transaction_id": "6075163",
		  "internal_id": 1354,
		  "formula_numeric": 112,
		  "formula_text": "A",
		  "subsidiary_id": null,
		  "subsidiary_name": null,
		  "type": null,
		  "creation_date": "2022-08-29"
		}
	   ]`

	dfEntSal := dataframe.ReadJSON(strings.NewReader(dataEntSal))
	dfDataLine := dataframe.ReadJSON(strings.NewReader(dataLine2))
	dfDataMove := dataframe.ReadJSON(strings.NewReader(dataMov))

	indexDelivery := 1
	indexLinea := 9

	//Concatena id con linea
	s := dfEntSal.Rapply(func(s series.Series) series.Series {

		deli := s.Elem(indexDelivery).String()
		linea := s.Elem(indexLinea).String()
		return series.Strings(deli + "-" + linea)
	})
	dfEntSal = dfEntSal.CBind(s.Rename("idLinea", "X0"))

	dfDataLine = dfDataLine.Mutate(dfDataLine.Col("Formula")).
		Rename("idLinea", "Formula")

	dfJoin := dfEntSal.InnerJoin(dfDataLine, "idLinea")

	indexItem := 19
	indexIden := 8

	a := dfJoin.Rapply(func(s series.Series) series.Series {

		item := s.Elem(indexItem).String()
		iden := s.Elem(indexIden).String()
		return series.Strings(iden + "-" + item)
	})
	dfJoin = dfJoin.CBind(a.Rename("idBusqueda", "X0"))

	//Comenzamos con movem

	//Concatena custbody_val_transaction_id con item_internal_id

	indexTrans := 1
	indexInt := 6

	c := dfDataMove.Rapply(func(s series.Series) series.Series {

		trans := s.Elem(indexTrans).String()
		int := s.Elem(indexInt).String()
		return series.Strings(trans + "-" + int)
	})
	dfDataMove = dfDataMove.CBind(c.Rename("idBusqueda", "X0"))

	dfFinal := dfJoin.InnerJoin(dfDataMove, "idBusqueda")

	indexCant := 2
	indexForm := 23

	n := dfFinal.Rapply(func(s series.Series) series.Series {
		cant, err := s.Elem(indexCant).Int()
		if err != nil {
			return series.Ints("NAN")
		}
		form, err := s.Elem(indexForm).Int()
		if err != nil {
			return series.Ints("NAN")
		}
		return series.Ints(form - cant)
	})
	dfFinal = dfFinal.CBind(n.Rename("diff", "X0"))

	return struct{}{}
}
