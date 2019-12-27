package dao

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/xavier268/go-demo-gin-gorm/internal/pkg/models"

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
var onceInternalDao sync.Once

// getDAOMemory provides a singleton DAO.
// It is lazily initialized and thread safe,
// using sync.once to ensure singleton pattern is enforced.
func getDAOMemory() *DAO {

	// define initMemoryDB to open and initialize a new in-memory database.
	// make it local, so that no one else can call it.
	// You HAVE TO call Close() when finished with it.
	initMemoryDB := func() {
		fmt.Println("Initializing DAO with in-memory sqlite3")
		db, err := gorm.Open("sqlite3", ":memory:")
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&models.Product{})
		internalDaoSingleton = &DAO{db}
	}

	// Execute only once during the program life.
	onceInternalDao.Do(initMemoryDB)

	return internalDaoSingleton
}

// getDAOPostgres opens and return a postgres DAO.
func getDAOPostgres() *DAO {
	// define initMemoryDB to open and initialize a new in-memory database.
	// make it local, so that no one else can call it.
	// You HAVE TO call Close() when finished with it.
	initMemoryDB := func() {
		fmt.Println("Initializing DAO with Postgres SQL")
		db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5432 user=postgres dbname=postgres password=secret")
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&models.Product{})
		internalDaoSingleton = &DAO{db}
	}

	// Execute only once during the program life.
	onceInternalDao.Do(initMemoryDB)

	return internalDaoSingleton
}

// GetDAO returns the postgresDAO by default.
func GetDAO() *DAO {
	// Todo : detect test flag, and use splite3/memory for tests and postgres otherwise
	return getDAOPostgres()
}

// Close the underlying database.
func (d *DAO) Close() error {
	fmt.Println("Closing database")
	return d.DB.Close()
}
