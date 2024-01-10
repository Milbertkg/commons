package sql

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb" // Controlador SQL Server
)

// SQLServerDB establece la conexión con SQL Server
func SQLServerDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlserver", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Conexión exitosa a la base de datos SQL Server.")
	return db, nil
}

// CheckSQLServerConnection realiza un Ping a la base de datos
func CheckSQLServerConnection(db *sql.DB) bool {
	err := db.Ping()
	if err == nil {
		return true
	}
	return false
}
