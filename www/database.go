package www

import (
	"go-server/database"
)

func RunDatabaseServer() {
	database.DatabaseConnect()
}
