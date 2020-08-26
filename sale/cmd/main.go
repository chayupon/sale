package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/chayupon/sale/service/v1/handle"
	_ "github.com/lib/pq"
)
func main() {
	//connect database
	const (
		dbHost     = "localhost"
		dbName     = "BotApp"
		dbUser     = "postgres"
		dbPassword = "tonkla727426"
		dbPort     = 5432
	)
	dbSale := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbSale)
	if err != nil {
		log.Fatal("Connect Fail")
	}
	log.Println("Connect")
	defer db.Close()
	handle.Router(db)
}
