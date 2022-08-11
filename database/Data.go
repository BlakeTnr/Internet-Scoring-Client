package database

import (
	"embed"
	"encoding/json"
)

//go:embed database.json
var database embed.FS

type WebsiteDatabase struct {
}

func (websiteDatabase WebsiteDatabase) ReadDatabase() string {
	bytes, _ := database.ReadFile("database.json")
	databaseString := string(bytes)
	return databaseString
}

type DatabaseWebsite struct {
	Domain string
}

func (websiteDatabase WebsiteDatabase) ParseDatabase() []DatabaseWebsite {
	database := websiteDatabase.ReadDatabase()
	var websites []DatabaseWebsite
	json.Unmarshal([]byte(database), &websites)
	return websites
}

func (websiteDatabase WebsiteDatabase) GetDomainByIndex(index int) string {
	websites := websiteDatabase.ParseDatabase()
	website := websites[index]
	return website.Domain
}

func (websiteDatabase WebsiteDatabase) GetNumberOfWebsites() int {
	websites := websiteDatabase.ParseDatabase()
	return len(websites)
}
