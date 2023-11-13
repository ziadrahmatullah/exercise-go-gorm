package crud

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/models"
	"gorm.io/gorm"
)

func UpdateAllWithNotNullDOBPets(db *gorm.DB){
	var pets []models.Pet
	db.Where("date_of_birth IS NOT NULL").Find(&pets)

	for _, pet := range pets {
		if pet.DateOfBirth != nil {
			*pet.DateOfBirth = pet.DateOfBirth.AddDate(1, 0, 0)
		}
	}
	result := db.Save(&pets)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
}