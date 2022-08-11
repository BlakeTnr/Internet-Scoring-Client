package encryption

import (
	"isc/database"
)

func GetKeyDatabase() IKeyDatabase {
	return database.KeyDatabase{}
}
