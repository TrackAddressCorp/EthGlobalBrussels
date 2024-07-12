package db

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"gorm.io/gorm"
)

func AddVote(id uint, signature string) error {
	err := DB.First(&models.Petition{}, id).UpdateColumn("votes", gorm.Expr("votes + ?", 1)).Error
	if err != nil {
		return err
	}
	err = DB.Create(&models.Vote{PetitionID: id, Signature: signature}).Error
	return err
}
