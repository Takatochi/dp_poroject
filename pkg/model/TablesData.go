package model

type Internal struct {
	ID      int
	Columns map[string]interface{}
}

type DataModelTables struct {
	TableName string     `json:"tableName"`
	Internal  []Internal `json:"internal"`
}
type DataModelTypeTables struct {
	TableName string   `json:"tableName"`
	Var       Internal `json:"var"`
}
