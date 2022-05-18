package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

/*

# Usage

## Using transaction

	// Begin transaction
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()

	// Using transaction
	tx.DB().Read( . . . )

	// Commit transaction
	tx.Commit()

## Using without transaction

	db.Read( . . . )

*/

// GormDB is gorm.DB of gorm package.
type GormDB = gorm.DB

// Errors definition.
var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

// Constants definition.
const (
	DriverMySQL = "mysql"
)

// Connect connects to BD.
func Connect(conf *Config) (*DB, error) {
	db, err := gorm.Open(conf.Driver, conf.ConnectionString())
	if err != nil {
		return nil, err
	}

	db.LogMode(conf.LogDebug)

	// Config
	if conf.MaxOpenConnections != 0 {
		db.DB().SetMaxOpenConns(conf.MaxOpenConnections)
	}

	if conf.MaxIdleConnections != 0 {
		db.DB().SetMaxIdleConns(conf.MaxIdleConnections)
	}

	if conf.ConnectionMaxLifetime != 0 {
		db.DB().SetConnMaxLifetime(time.Duration(conf.ConnectionMaxLifetime) * time.Second)
	}

	db.SingularTable(true)

	return &DB{db}, nil
}

// DB is wrapper of gorm.DB.
type DB struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) *DB {
	return &DB{db: db}
}

// DB returns current instance of *gorm.DB.
func (d *DB) DB() *GormDB {
	return d.db
}

// Begin opens a transaction.
func (d *DB) Begin() *DB {
	return &DB{d.db.Begin()}
}

// RollbackUnlessCommitted rollbacks if a transaction not committed.
func (d *DB) RollbackUnlessCommitted() {
	d.db.RollbackUnlessCommitted()
}

// Commit closes and saves a DB transaction.
func (d *DB) Commit() *gorm.DB {
	return d.db.Commit()
}

// Config contains connection info of DB.
type Config struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int64
	Database string
	LogDebug bool

	MaxIdleConnections    int
	MaxOpenConnections    int
	ConnectionMaxLifetime int
}

// ConnectionString returns MySQL connection string
func (d *Config) ConnectionString() string {
	switch d.Driver {
	case DriverMySQL:
		return fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8mb4&loc=Local",
			d.Username,
			d.Password,
			d.Host,
			d.Port,
			d.Database,
		)
	}
	return ""
}
