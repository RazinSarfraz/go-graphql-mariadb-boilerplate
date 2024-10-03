package mariadb

import (
	"fmt"
	"sync"

	"github.com/appointment-api/conf"
	"github.com/appointment-api/graph/model"
	"github.com/appointment-api/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gdb *gorm.DB
var storeOnce sync.Once
var store repository.Store

type Store struct {
	db *gorm.DB
}

func SharedStore() repository.Store {
	storeOnce.Do(func() {
		err := initDb()
		if err != nil {
			panic(err)
		}
		store = NewStore(gdb)
	})
	return store
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func initDb() error {
	cfg := conf.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MariaDbDataSource.User, cfg.MariaDbDataSource.Password, cfg.MariaDbDataSource.Addr, cfg.MariaDbDataSource.Port, cfg.MariaDbDataSource.Database)

	var err error
	gdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Connection pool settings
	sqlDB, err := gdb.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)

	if cfg.MariaDbDataSource.EnableAutoMigrate {
		var tables = []interface{}{
			&model.JobListing{},
		}
		for _, table := range tables {
			if err := gdb.AutoMigrate(table); err != nil {
				return err
			}
		}
	}

	return nil
}
