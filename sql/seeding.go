package sql

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/crud"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/models"
	"gorm.io/gorm"
)

func Seeding(db *gorm.DB){
	crud.CreateCategory(db, "Dog")
	crud.CreateCategory(db, "Cat")
	var owners = []models.Owner{
		{Name: "Andy", Address: "Apple Street No.13"},
		{Name: "Bob", Address: "Mango Street No.23"},
		{Name: "Celine", Address: "Grape Street No.25"},
		{Name: "David", Address: "Pear Street No.29"},
		{Name: "Evelyn", Address: "Banana Street No.2"},
		{Name: "Brynn", Address: "Coconut Street No.2"},
	}
	crud.CreateOwners(db, owners)

	var pets = []models.Pet{
		{
			Name:          "Coco",
			PetCategoryId: 1,
			IsAggresive:   false},
		{
			Name:          "Jack",
			PetCategoryId: 1,
			IsAggresive:   true,
			DateOfBirth:   Date(2022, 1, 1)},
		{
			Name:          "Max",
			PetCategoryId: 1,
			IsAggresive:   true,
			DateOfBirth:   Date(2023, 1, 2)},
		{
			Name:          "Buddy",
			PetCategoryId: 2,
			IsAggresive:   true,
			DateOfBirth:   Date(2019, 5, 1)},
		{
			Name:          "Kitty",
			PetCategoryId: 2,
			IsAggresive:   true},
	}

	crud.CreatePets(db, pets)

	var ownerPets = []models.OwnerPet{
		{OwnerId: 1, PetId: 5, StartDate: Date(2023, 1, 1)},
		{OwnerId: 1, PetId: 4, StartDate: Date(2023, 1, 2)},
		{OwnerId: 4, PetId: 3, StartDate: Date(2023, 1, 3)},
		{OwnerId: 3, PetId: 1, StartDate: Date(2023, 1, 4)},
	}
	crud.CreateOwnerPets(db, ownerPets)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}