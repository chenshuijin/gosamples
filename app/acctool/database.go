package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbs = map[string]*gorm.DB{}

func InitDatabaseConfig(name string, c DBConf) error {
	db, err := gorm.Open(c.DriverName, c.DataSource)
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(c.MaxIdleConns)
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	dbs[name] = db
	return nil
}

func DefaultDB() *gorm.DB {
	return dbs["default"]
}
