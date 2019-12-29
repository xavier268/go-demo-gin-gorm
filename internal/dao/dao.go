package dao

import (
	"errors"
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

// Source is used to generate a DAO object (singleton)
type Source struct {
	dao            *DAO
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

// GetDAO returns the postgresDAO by default.
func (s *Source) GetDAO() *DAO {

	s.once.Do(func() {
		fmt.Printf("Initializing database with %s (%s)\n", s.driver, s.params)
		db, err := gorm.Open(s.driver, s.params)
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&models.Product{})
		s.dao = &DAO{db}
	})
	return s.dao
}

// Close the underlying database.
func (d *DAO) Close() error {
	if d.DB != nil {
		fmt.Println("Closing database")
		return d.DB.Close()
	}
	return errors.New("Cannot close a non existing DAO")
}

// Close the underlying database
func (s *Source) Close() error {
	return s.dao.Close()
}
