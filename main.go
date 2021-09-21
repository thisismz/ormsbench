package main

import (
	"context"
	"database/sql"
	models "ormsbench/my_models"

	"fmt"
	"log"

	_ "github.com/lib/pq"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Orders struct {
	OrderID    string
	CustomerID string
}

var orders Orders

func main() {
	//define connections
	dsn := "host=localhost user=admin dbname=northwind port=5432 sslmode=disable"
	dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbBioler, err := sql.Open("postgres", `dbname=northwind host=localhost user = admin`)
	err = dbBioler.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer dbBioler.Close()

	// define queries
	dbGorm.First(&orders)
	fmt.Println(orders.CustomerID)

	got, err := models.Orders().One(context.Background(), dbBioler)
	fmt.Println(got.CustomerID)
	got1, err := models.Orders().All(context.Background(), dbBioler)
	for _, v := range got1 {
		fmt.Println(v.CustomerID)
	}
}
