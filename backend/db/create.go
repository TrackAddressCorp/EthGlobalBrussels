package db

import "github.com/TrackAddressCorp/EthGlobalBrussels/models"

func CreatePetition(petition *models.Petition) error {
	return DB.Create(&petition).Error
}
