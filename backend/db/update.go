package db

import (
	"errors"

	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"gorm.io/gorm"
)

func AddSign(id uint, NullifierHash string) error {
	var existingSign models.Sign
	result := DB.Where("nullifier_hash = ?", NullifierHash).First(&existingSign)

	if result.Error == nil {
		return errors.New("sign already exists")
	} else if result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	err := DB.Create(&models.Sign{PetitionID: id, NullifierHash: NullifierHash}).Error
	if err != nil {
		return err
	}

	err = DB.Model(&models.Petition{}).Where("id = ?", id).UpdateColumn("signs", gorm.Expr("signs + ?", 1)).Error
	return err
}
