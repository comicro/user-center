package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"user-center/conf"
)

var db *DB

type DB struct {
	orm *gorm.DB
}

// load database
func Load() {
	orm, err := newORM()
	if err != nil {
		panic(err)
	}
	db = new(DB)
	db.orm = orm
}

func newORM() (*gorm.DB, error) {
	c := conf.GetDataBaseConfig()
	url := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Host, c.Username, c.Database, c.Password,
	)
	return gorm.Open("postgres", url)
}
