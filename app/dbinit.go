package app

import (
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

// DBconn manages the db connection state
type DBconn struct {
	DB *gorm.DB
}

var onceDB sync.Once
var dbinst *DBconn

// GetDBInstance will get us the DB connection using singleton pattern
func GetDBInstance() *DBconn {
	onceDB.Do(func() {
		d, err := initPostgresDB()
		if err != nil {
			panic(err)
		}
		dbinst = &DBconn{DB: d}
	})
	return dbinst
}

func initDb() *gorm.DB {
	// check file, if exists
	if _, err := os.Stat("./auth.sqlite3"); os.IsNotExist(err) {
		panic("DB doesnt exists")
	}
	db, err := gorm.Open("sqlite3", "./auth.sqlite3")
	// Error
	if err != nil {
		panic(err)
	}
	// Display SQL queries
	db.LogMode(true)
	return db
}

// initPostgresDB is to create a new GORM connection with postgres
func initPostgresDB() (*gorm.DB, error) {
	conf := GetConfig()
	var db *gorm.DB
	// If conf.Type is empty. Expected value postgresql
	if len(conf.Type) == 0 {
		return db, errors.New("conf.Type is empty")
	}
	// Validate other params
	if conf.Host == "" || conf.User == "" || conf.Name == "" {
		return db, errors.New("ConnectionString is invalid")
	}
	// SSL mode check for postgres
	if conf.SSLMode == "" {
		return db, errors.New("Postgress SSL mode not configured. Please add this to .env or env variable : PGSSLMODE=disable/enable")
	}
	if conf.Type == "postgresql" && (conf.SSLMode == "disable" || conf.SSLMode == "require") {
		db, err := gorm.Open("postgres", "host="+conf.Host+" port="+conf.Port+" user="+conf.User+" dbname="+conf.Name+" sslmode="+conf.SSLMode+" password="+conf.Password)
		if err != nil {
			return db, err
		}
		return db, err
	}
	return db, errors.Errorf("%v database is not supported(Type:Postgresql) or SSL mode not confiugured(Please add this to .env or env varibale : PGSSLMODE=disable/enable)", conf.Type)
}
