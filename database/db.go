package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/matkinhig/developer-community/config"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVE, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
