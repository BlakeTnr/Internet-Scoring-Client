package database

import (
	"embed"
	"encoding/json"
)

//go:embed database.json
var database embed.FS

func ReadDatabase() string {
	bytes, _ := database.ReadFile("database.json")
	databaseString := string(bytes)
	return databaseString
}

type DatabaseWebsite struct {
	Domain string
}

func ParseDatabase() []DatabaseWebsite {
	database := ReadDatabase()
	var websites []DatabaseWebsite
	json.Unmarshal([]byte(database), &websites)
	return websites
}

func GetDomainByIndex(index int) string {
	websites := ParseDatabase()
	website := websites[index]
	return website.Domain
}

func GetNumberOfWebsites() int {
	websites := ParseDatabase()
	return len(websites)
}
