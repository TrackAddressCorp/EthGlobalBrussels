package db

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("smartpetition.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&models.Petition{}, &models.Vote{})
	return err
}
