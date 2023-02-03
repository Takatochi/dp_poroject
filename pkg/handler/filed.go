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
		return
	}
	listTables, err := showTables(bd)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "server not found connect with bd"})
		logger.Infof("server not found connect with bd %d", port)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": listTables, "info": "File uploaded successfully", "list": allTables})

}

func retrieveAllDataFromAllTables(db *sql.DB) ([]model.DataModelTables, error) {
	tableName, err := showTables(db)
	if err != nil {
		return nil, err
	}
	var DataModelTablesList []model.DataModelTables
	//&model.Server{}
	var DataModelTables model.DataModelTables
	for index, element := range tableName {

		fmt.Print(index, element)
		//	// Retrieve all data from each table
		tableRows, err := db.Query("SELECT * FROM " + string(element))
		if err != nil {
			return nil, err
		}
		defer tableRows.Close()
		//	// Print the data from each table
		fmt.Println("Data from table:", element)
		DataModelTables.TableName = string(element)
		for tableRows.Next() {
			var columns []string
			var columnPointers []interface{}

			// Create a slice of pointers to the columns
			columns, err = tableRows.Columns()
			if err != nil {

				return nil, err
			}
			for range columns {
				columnPointers = append(columnPointers, new(string))
			}

			// Scan the columns into the slice of pointers
			if err := tableRows.Scan(columnPointers...); err != nil {
				fmt.Println(err)
				return nil, err
			}

			// Print the column values
			//fmt.Println(columns)
			var data []model.Internal
			var сolumnsBL = make(map[string]interface{})
			for i, column := range columns {
				сolumnsBL[column] = *columnPointers[i].(*string)
				//data = append(data, *columnPointers[i].(*string))
				//fmt.Println(column, *columnPointers[i].(*string))

			}
			data = append(data, model.Internal{ID: 0, Columns: сolumnsBL})
			DataModelTables.Internal = data
			//data = append(data, model.Internal{ID: i, Columns: map[string]interface{}{column}})
			// end

		}
		DataModelTablesList = append(DataModelTablesList, DataModelTables)
	}
	fmt.Println(DataModelTablesList)
	for _, d := range DataModelTablesList {
		fmt.Println(d.Internal)
	}
	return DataModelTablesList, nil
}
