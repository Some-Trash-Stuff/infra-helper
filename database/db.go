package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Some-Trash-Stuff/infra-helper/configs"
	_ "github.com/microsoft/go-mssqldb"
)

// NewSqlConnection cria uma nova conexão com o banco de dados
func NewSqlConnection(database string) (*sql.DB, error) {

	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		configs.Configs.Database.Server, configs.Configs.Database.User, configs.Configs.Database.Password, configs.Configs.Database.Port, database)

	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Connected!")
	return db, nil
}
