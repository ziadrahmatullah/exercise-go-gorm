package main

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-gorm/sql"
)

func main() {
	db := database.ConnectDB()
	sql.Seeding(db)
	
}
