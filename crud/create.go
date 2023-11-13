package crud

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/models"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, name string) {
	petCategory := &models.PetCategory{Name: name}
	result := db.Select("Name").Create(&petCategory)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
}

func CreateOwners(db *gorm.DB, owners []models.Owner) {
	result := db.Create(&owners)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)

}

func CreatePets(db *gorm.DB, pets []models.Pet) {
	result := db.Create(&pets)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
}

func CreateOwnerPets(db *gorm.DB, ownerPets []models.OwnerPet){
	result := db.Create(&ownerPets)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
}
