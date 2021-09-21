package main

import (
	"context"
	"database/sql"
	"log"
	models "ormsbench/my_models"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BenchmarkGromSelectInts1(b *testing.B) {
	benchmarkGromSelectInts(b)
}
func BenchmarkBoilerSelectInts1(b *testing.B) {
	benchmarkBoilerSelectInts(b)
}
func BenchmarkGromSelectAll(b *testing.B) {
	benchmarkGromSelectAll(b)
}
func BenchmarkBoilerSelectAll(b *testing.B) {
	benchmarkBoilerSelectAll(b)
}
func BenchmarkGromSelectAll10(b *testing.B) {
	benchmarkGromSelectAllLimit(b, 10)
}
func BenchmarkBoilerSelectAll10(b *testing.B) {
	benchmarkBoilerSelectAllLimit(b, 10)
}
func BenchmarkGromSelectAll100(b *testing.B) {
	benchmarkGromSelectAllLimit(b, 100)
}
func BenchmarkBoilerSelectAll100(b *testing.B) {
	benchmarkBoilerSelectAllLimit(b, 100)
}
func BenchmarkGromSelectAll1000(b *testing.B) {
	benchmarkGromSelectAllLimit(b, 1000)
}
func BenchmarkBoilerSelectAll1000(b *testing.B) {
	benchmarkBoilerSelectAllLimit(b, 1000)
}

// functions implementation
func benchmarkGromSelectInts(b *testing.B) {
	db := postgressGromConn()
	b.ResetTimer()
	type Orders struct {
		OrderID    string
		CustomerID string
	}
	var orders Orders

	for n := 0; n < b.N; n++ {
		db.First(&orders)
	}
}
func benchmarkBoilerSelectInts(b *testing.B) {
	db := postgresBoilerConn()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := models.Orders().One(context.Background(), db)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func benchmarkGromSelectAll(b *testing.B) {
	db := postgressGromConn()
	b.ResetTimer()
	type Orders struct {
		OrderID    string
		CustomerID string
	}
	var orders Orders

	for n := 0; n < b.N; n++ {
		db.Find(&orders)
	}
}
func benchmarkBoilerSelectAll(b *testing.B) {
	db := postgresBoilerConn()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := models.Orders().All(context.Background(), db)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// select with limit
func benchmarkGromSelectAllLimit(b *testing.B, limit int) {
	db := postgressGromConn()
	b.ResetTimer()
	type Orders struct {
		OrderID    string
		CustomerID string
	}
	var orders Orders

	for n := 0; n < b.N; n++ {
		db.Limit(limit).Find(&orders)
	}
}
func benchmarkBoilerSelectAllLimit(b *testing.B, limit int) {
	db := postgresBoilerConn()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := models.Orders(qm.Limit(limit)).All(context.Background(), db)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//
// Connection helpers
//

func postgressGromConn() *gorm.DB {
	dsn := "host=localhost user=admin dbname=northwind port=5432 sslmode=disable"
	dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return dbGorm
}
func postgresBoilerConn() *sql.DB {
	dbBioler, err := sql.Open("postgres", `dbname=northwind host=localhost user = admin`)
	err = dbBioler.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return dbBioler
}
