package internetchecker

import (
	database "isc/database"
)

func GetWebsiteDatabase() IWebsiteDatabase {
	return database.WebsiteDatabase{}
}
