package dao

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/xavier268/go-demo-gin-gorm/internal/models"

	// import sqlite3 dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// import postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DAO is the main Data Access Object.
type DAO struct {
	*gorm.DB
}

// internalDaoSingleton is a package var, lazily initiated, as a singleton.
// Avoid accessing it directly, use GetDAO() instead.
var internalDaoSingleton *DAO
var internalDaoOnce sync.Once

// driver and associated parameters.
var internalDriver, internalParams string

// UseDriverMemory use memory database
func UseDriverMemory() {
	if internalDaoSingleton != nil {
		panic("You cannot set the driver type once a database has been created")
	}
	internalDriver = "sqlite3"
	internalParams = ":memory:"
}

// UseDriverPostgres use postgres DB
func UseDriverPostgres() {
	if internalDaoSingleton != nil {
		panic("You cannot set the driver type once a database has been created")
	}
	internalDriver = "postgres"
	internalParams = "host=localhost sslmode=disable port=5432 user=postgres dbname=postgres password=secret"
}

// GetDAO returns the postgresDAO by default.
func GetDAO() *DAO {
	if internalDriver == "" {
		UseDriverPostgres()
	}

	internalDaoOnce.Do(func() {
		fmt.Printf("Initializing database with %s (%s)\n", internalDriver, internalParams)
		db, err := gorm.Open(internalDriver, internalParams)
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&models.Product{})
		internalDaoSingleton = &DAO{db}
	})

	return internalDaoSingleton
}

// Close the underlying database.
func (d *DAO) Close() error {
	fmt.Println("Closing database")
	return d.DB.Close()
}
