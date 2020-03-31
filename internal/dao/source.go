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

// Source is used to generate a DAO object (singleton)
type Source struct {
	db             *gorm.DB
	driver, params string
	once           sync.Once
}

// NewMemorySource generate a  memory database source
func NewMemorySource() *Source {
	s := new(Source)
	s.driver, s.params = "sqlite3", ":memory:"
	return s
}

// NewPostgresSource generates a postgres source
func NewPostgresSource() *Source {
	s := new(Source)
	s.driver, s.params = "postgres",
		"host=localhost sslmode=disable port=5432 user=postgres dbname=postgres password=secret"
	return s
}

// GetDB (lazily open and) returns the database associated with that source
func (s *Source) GetDB() *gorm.DB {

	s.once.Do(func() {
		fmt.Printf("Initializing database with %s (%s)\n", s.driver, s.params)
		db, err := gorm.Open(s.driver, s.params)
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&models.Product{})
		s.db = db
	})
	return s.db
}

// Close the underlying database
func (s *Source) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
