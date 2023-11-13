package crud

import (
	"fmt"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/models"
	"gorm.io/gorm"
)

func GetFirstOwnerData(db *gorm.DB) (string, error){
	var owner models.Owner
	result:=db.First(&owner)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)	
	return ToJson(owner)
}

func GetOwnerDataByAddress(db *gorm.DB, address string) (string, error){
	var owner []models.Owner
	result:=db.Find(&owner, "address = ?", address)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)	
	return ToJson(owner)
}

func GetOwnerDataByContains(db *gorm.DB, contain string) (string, error){
	var owner []models.Owner
	contain = "%" + contain + "%"
	result:=db.Find(&owner, "name ilike ?", contain)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)	
	return ToJson(owner)
}

func GetOwnerDataNotContainsVowel(db *gorm.DB) (string, error){
	var owner []models.Owner
	result:=db.Where("name NOT SIMILAR TO '%(a|i|u|e|o|A|I|U|E|O)%'").Find(&owner)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)	
	return ToJson(owner)
}

func GetOwnerDataByTriId(db *gorm.DB) (string, error){
	var owner []models.Owner
	result:=db.Find(&owner, "owner_id = ? or owner_id = ? or owner_id = ?", 1,2,3)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)	
	return ToJson(owner)
}


func GetPetsNotHaveDateBirth(db *gorm.DB)(string, error){
	var pets []models.Pet
	result:= db.Find(&pets, "date_of_birth IS NULL").Scan(&pets)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
	return ToJson(pets)
}

func GetPetsAtleasTwoYears(db *gorm.DB)(string, error){
	var pets []models.Pet
	currentDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	twoYearsAgo := currentDate.AddDate(-2, 0, 0)
	result := db.Where("date_of_birth <= ?", twoYearsAgo).Find(&pets)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
	return ToJson(pets)
}

func GetPetsNeverHadOwner(db *gorm.DB) (string, error){//-----
	var pets []models.Pet
	result := db.Joins("LEFT JOIN owner_pets ON pets.pet_id = owner_pets.pet_id").
		Where("owner_pets.pet_id IS NULL").
		Find(&pets)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
	return ToJson(pets)
}

func GetOwnersWithPets(db *gorm.DB) (string, error){//---
	var owners []models.Owner
	result := db.Preload("Pets").Find(&owners)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
	return ToJson(owners)
}

func GetPetCategories(db *gorm.DB) (string, error){//----
	var petCategory models.PetCategory
	result := db.Model(&models.Pet{}).
		Select("pet_categories.pet_category_id, pet_categories.name, COUNT(pets.pet_id) as pet_count").
		Joins("LEFT JOIN pet_categories ON pets.pet_category_id = pet_categories.pet_category_id").
		Group("pet_categories.pet_category_id").
		Order("pet_count").
		Limit(1).
		Scan(&petCategory)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
	return ToJson(petCategory)
}

func GetTwoOwnerDataOrderById(db *gorm.DB) (string, error){
	var owners []models.Owner
	result := db.Order("owner_id").Limit(2).Find(&owners)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
	return ToJson(owners)
}

func GetSecondOwnerDataOrderById(db *gorm.DB) (string, error){
	var owners []models.Owner
	result := db.Order("owner_id").Offset(2).Limit(2).Find(&owners)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
	return ToJson(owners)
}

func GetOwnersNameAndTheirPetCount(db *gorm.DB)(string, error){
	var ownerPetCounts []entity.OwnerPetCount
	query := `
		SELECT owners.owner_id, owners.name, COUNT(owner_pets.pet_id) as pet_count
		FROM owners
		LEFT JOIN owner_pets ON owners.owner_id = owner_pets.owner_id
		GROUP BY owners.owner_id, owners.name
		ORDER BY owners.owner_id
	`
	result:= db.Raw(query).Scan(&ownerPetCounts)
	fmt.Printf("rowsAfected : %d\n", result.RowsAffected)
	return ToJson(ownerPetCounts)
}

