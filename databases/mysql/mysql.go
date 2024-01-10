package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Controlador MySQL
)

// MySQLDB establece la conexión con MySQL
func MySQLDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Conexión exitosa a la base de datos MySQL.")
	return db, nil
}

// CheckMySQLConnection realiza un Ping a la base de datos
func CheckMySQLConnection(db *sql.DB) bool {
	err := db.Ping()
	if err == nil {
		return true
	}
	return false
}
