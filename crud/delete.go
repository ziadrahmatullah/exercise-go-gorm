package crud

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/models"
	"gorm.io/gorm"
)

func DeletePetWhichNeverHaveOwner(db *gorm.DB){
	result := db.Where("pet_id NOT IN (SELECT pets.pet_id FROM pets LEFT JOIN owner_pets ON pets.pet_id = owner_pets.pet_id WHERE owner_pets.pet_id IS NOT NULL)").Delete(&models.Pet{})
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
}
