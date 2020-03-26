package database

import (
	"github.com/jinzhu/gorm"
	"github.com/matkinhig/developer-community.git/config"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVE, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
