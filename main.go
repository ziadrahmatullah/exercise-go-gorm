package main

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/crud"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/database"
)

func main() {
	db := database.ConnectDB()

	// You can change the function bellow
	// Uncoment For seeding
	// sql.Seeding(db)

	// Uncoment for try Read
	idx, err:= crud.GetOwnersNameAndTheirPetCount(db)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(idx)

	// Uncoment for try Delete and Update
	// crud.DeletePetWhichNeverHaveOwner(db)
	// crud.UpdateAllWithNotNullDOBPets(db)

}
