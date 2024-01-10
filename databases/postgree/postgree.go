package postgree

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Controlador PostgreSQL
)

// PostgreSQLDB establece la conexión con PostgreSQL
func PostgreSQLDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Conexión exitosa a la base de datos PostgreSQL.")
	return db, nil
}

// CheckPostgreSQLConnection realiza un Ping a la base de datos
func CheckPostgreSQLConnection(db *sql.DB) bool {
	err := db.Ping()
	if err == nil {
		return true
	}
	return false
}
