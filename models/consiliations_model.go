package models

type AjuTrans struct {
	TIPO_MOV       string
	TRANSACTION_ID string
	INTERNAL_ID    string
	FROM_QUALITY   string
	TO_QUALITY     string
	CANTIDAD       string
}

type EntSal struct {
	DELIVERY          string
	DOCUMENTO         string
	TIPO_DOC          string
	IDENTIFICADOR_MOV string
	MONEDA            string
	LINEA             string
	CANTIDAD          string
	NO_PEDIMENTO      string
	FECHA_PEDIMENTO   string
	NOMBRE_ADUANA     string
	RETORNO_BODEGA    string
	ESTADO_ALMACEN    string
}

type Line2 struct {
	Internalid string
	Formula    string
	Item       string
}

type Movimientos struct {
	recordType     string
	Internalid     string
	formulanumeric string
}
