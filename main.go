package main

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/crud"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/database"
)

func main() {
	db := database.ConnectDB()
	// sql.Seeding(db)
	idx, err:= crud.GetOwnersNameAndTheirPetCount(db)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(idx)
	// crud.DeletePetWhichNeverHaveOwner(db)
	// crud.UpdateAllWithNotNullDOBPets(db)

}
