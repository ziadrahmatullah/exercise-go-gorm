package sql

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/crud"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/models"
	"gorm.io/gorm"
)

func Seeding(db *gorm.DB) {
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
			IsAggresive:   false,
			DateOfBirth:   nil},
		{
			Name:          "Jack",
			PetCategoryId: 1,
			IsAggresive:   true,
			DateOfBirth:   toDate("2022-01-01")},
		{
			Name:          "Max",
			PetCategoryId: 1,
			IsAggresive:   true,
			DateOfBirth:   toDate("2023-01-02)")},
		{
			Name:          "Buddy",
			PetCategoryId: 2,
			IsAggresive:   true,
			DateOfBirth:   toDate("2019-05-01")},
		{
			Name:          "Kitty",
			PetCategoryId: 2,
			IsAggresive:   true,
			DateOfBirth:   nil},
	}

	crud.CreatePets(db, pets)

	var ownerPets = []models.OwnerPet{
		{OwnerId: 1, PetId: 5, StartDate: toDate("2023-01-01"), EndDate: nil},
		{OwnerId: 1, PetId: 4, StartDate: toDate("2023-01-02"), EndDate: nil},
		{OwnerId: 4, PetId: 3, StartDate: toDate("2023-01-03"), EndDate: nil},
		{OwnerId: 3, PetId: 1, StartDate: toDate("2023-01-04"), EndDate: nil},
	}
	crud.CreateOwnerPets(db, ownerPets)
}

func toDate(dateString string) *time.Time {
	parsedDate, _ := time.Parse("2006-01-02", dateString)
	return &parsedDate
}
