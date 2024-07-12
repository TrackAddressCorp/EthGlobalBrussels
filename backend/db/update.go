package db

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"gorm.io/gorm"
)

func AddSign(id uint, NullifierHash string) error {
	err := DB.First(&models.Petition{}, id).UpdateColumn("signs", gorm.Expr("votes + ?", 1)).Error
	if err != nil {
		return err
	}
	err = DB.Create(&models.Vote{PetitionID: id, NullifierHash: NullifierHash}).Error
	return err
}
