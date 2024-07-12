package db

import "github.com/TrackAddressCorp/EthGlobalBrussels/models"

func GetPetition(id uint) (models.Petition, error) {
	var petition models.Petition

	err := DB.First(&petition, id).Error

	return petition, err
}
