// Go connection Sample Code:
package main
import (
	_ "github.com/microsoft/go-mssqldb"
	"database/sql"
	"context"
	"log"
	"fmt"
	//"errors"
)

var db *sql.DB
var server = "svdb-social.database.windows.net"
var port = 1433
var user = "useradmin"
var password = "_k8A801_"
var database = "socialdb_data"

func main() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
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
	fmt.Scanln() // Esperar a que el usuario presione Enter
}