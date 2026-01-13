// Go connection Sample Code:
package migrations

import (
	"context"
	"fmt"
	db_helper "infra-helper/helper/database"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var database = "sqldb-dev-bitchens-one"

func main() {
	db, err := db_helper.NewSqlConnection(database)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")

}
