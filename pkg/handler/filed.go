package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"project/pkg/Database/VirtualSql"
	mysqldump "project/pkg/Database/dumpmysql"
	"project/pkg/handler/mapHadler"
	"project/pkg/logger"
	"project/pkg/model"
)

type tables string

func (h *Index) GetFileSqL(ctx *gin.Context) {

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	defer file.Close()
	port, err := h.getPortFromContextPOST(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Errorf("Failed to get port from %s", err.Error())
		return
	}

	if ServerTree.Empty() {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server tree is empty"})
		logger.Errorf("server tree is empty %d", port)
		return
	}

	server, found := ServerTree.Get(port)
	if server == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server is nil"})
		logger.Errorf("server is nil %d", port)
		return
	}
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "server not found"})
		logger.Infof("Server not found port %d", port)
		return
	}

	dir, err := saveFile("upload/", file, header)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("os create file err: %s", err.Error()))
		return
	}

	var store *VirtualSql.VirtualMySQLDatabase
	bd, err := openVirtualSql(store, server.(mapHadler.ListServerSql).Config)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("problem with connect : %s", err.Error()))
		return
	}
	err = loaderDataSql(dir, bd)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("problem with saving : %s", err.Error()))
		return
	}
	listTables, err := showTables(bd)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "server not found connect with bd"})
		logger.Infof("server not found connect with bd %d", port)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": listTables, "info": "File uploaded successfully"})
}

func saveFile(dir string, file multipart.File, header *multipart.FileHeader) (string, error) {
	// Create a new file
	dir += header.Filename
	f, err := os.Create(dir)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Copy the uploaded file to the created file
	io.Copy(f, file)

	return dir, nil
}

func openVirtualSql(db VirtualSql.Database, store *VirtualSql.ConfigVirtual) (*sql.DB, error) {

	DBStore, err := db.Open(store)
	if err != nil {
		return nil, err
	}
	return DBStore, nil
}
func loaderDataSql(dir string, db *sql.DB) error {
	err := mysqldump.Load(db, dir)
	if err != nil {
		return err
	}
	return nil
}
func showTables(db *sql.DB) ([]tables, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}

	var tablesArr []tables
	for rows.Next() {
		var name tables
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tablesArr = append(tablesArr, name)
	}
	return tablesArr, nil
}

func (h *Index) GetTableWITHPort(ctx *gin.Context) {

	port, err := h.getPortFromContextPATH(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Errorf("Failed to get port from %s", err.Error())
		return
	}
	if ServerTree.Empty() {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server tree is empty"})
		logger.Errorf("server tree is empty %d", port)
		return
	}

	server, found := ServerTree.Get(port)
	if server == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server is nil"})
		logger.Errorf("server is nil %d", port)
		return
	}
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "server not found"})
		logger.Infof("Server not found port %d", port)
		return
	}

	var store *VirtualSql.VirtualMySQLDatabase

	bd, err := openVirtualSql(store, server.(mapHadler.ListServerSql).Config)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("problem with connect : %s", err.Error()))
		return
	}
	allTables, err := retrieveAllDataFromAllTables(bd)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "server not found connect with bd"})
		logger.Infof("server not found connect with bd %d", port)
		return
	}
	//listTables, err := showTables(bd)
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, gin.H{"error": "server not found connect with bd"})
	//	logger.Infof("server not found connect with bd %d", port)
	//	return
	//}
	typeTables, err := retrieveVarTypeTables(bd)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "server not found connect with bd"})
		logger.Infof("server not found connect with bd %d", port)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"info": "File uploaded successfully", "type": typeTables, "data": allTables})

}

func retrieveAllDataFromAllTables(db *sql.DB) ([]model.DataModelTables, error) {
	tableName, err := showTables(db)
	if err != nil {
		return nil, err
	}
	var DataModelTablesList []model.DataModelTables
	var DataModelTables model.DataModelTables
	for _, element := range tableName {
		tableRows, err := db.Query("SELECT * FROM " + string(element))
		if err != nil {
			return nil, err
		}

		DataModelTables.TableName = string(element)
		for tableRows.Next() {
			var columns []string
			var columnPointers []interface{}

			columns, err = tableRows.Columns()
			if err != nil {
				return nil, err
			}

			// Get column types
			columnTypes, err := tableRows.ColumnTypes()
			if err != nil {
				return nil, err
			}

			// Create a slice of pointers to the columns based on their type
			for i := range columns {
				switch columnTypes[i].DatabaseTypeName() {
				case "BLOB":
					columnPointers = append(columnPointers, new([]byte))
				case "INT", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT":
					var in sql.NullInt64
					columnPointers = append(columnPointers, &in)
				case "FLOAT", "DOUBLE", "DECIMAL":
					var f sql.NullFloat64
					columnPointers = append(columnPointers, &f)
				case "CHAR", "VARCHAR", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT":
					var s sql.NullString
					columnPointers = append(columnPointers, &s)
				default:
					columnPointers = append(columnPointers, new(string))
				}
			}

			// Scan the columns into the slice of pointers
			if err := tableRows.Scan(columnPointers...); err != nil {
				return nil, err
			}

			var data []model.Internal
			var сolumnsBL = make(map[string]interface{})
			for i, column := range columns {
				switch columnTypes[i].DatabaseTypeName() {
				case "BLOB":
					сolumnsBL[column] = *columnPointers[i].(*[]byte)
				case "INT", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT":
					сolumnsBL[column] = columnPointers[i].(*sql.NullInt64).Int64
				case "FLOAT", "DOUBLE", "DECIMAL":
					сolumnsBL[column] = columnPointers[i].(*sql.NullFloat64).Float64
				case "CHAR", "VARCHAR", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT":
					сolumnsBL[column] = columnPointers[i].(*sql.NullString).String
				default:
					сolumnsBL[column] = *columnPointers[i].(*string)
				}
				data = append(data, model.Internal{ID: i, Columns: сolumnsBL})
			}

			DataModelTables.Internal = data
		}
		DataModelTablesList = append(DataModelTablesList, DataModelTables)
		err = tableRows.Close()
		if err != nil {
			return nil, err
		}
	}

	return DataModelTablesList, nil
}

func retrieveVarTypeTables(db *sql.DB) ([]model.DataModelTypeTables, error) {
	tableName, err := showTables(db)
	if err != nil {
		return nil, err
	}

	var DataModelTypeTablesList []model.DataModelTypeTables
	var DataModelTypeTables model.DataModelTypeTables

	for _, element := range tableName {
		tableRows, err := db.Query("SELECT * FROM " + string(element))
		if err != nil {
			return nil, err
		}

		DataModelTypeTables.TableName = string(element)
		for tableRows.Next() {
			var columns []string
			columns, err = tableRows.Columns()
			if err != nil {
				return nil, err
			}

			// Get column types
			columnTypes, err := tableRows.ColumnTypes()
			if err != nil {
				return nil, err
			}
			var сolumnsBL = make(map[string]interface{})
			//var data []model.Internal
			for i, column := range columns {
				сolumnsBL[column] = columnTypes[i].DatabaseTypeName()

			}
			//data = append(data, model.Internal{ID: -1, Columns: сolumnsBL})
			DataModelTypeTables.Var = model.Internal{ID: -1, Columns: сolumnsBL}
		}

		DataModelTypeTablesList = append(DataModelTypeTablesList, DataModelTypeTables)
		err = tableRows.Close()
		if err != nil {
			return nil, err
		}
	}

	//DataModelTypeTables.Var = model.Internal{ID: -1, Columns: сolumnsBL}
	return DataModelTypeTablesList, nil
}
