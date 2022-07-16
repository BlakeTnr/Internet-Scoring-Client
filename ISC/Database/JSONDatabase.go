package Database

import (
	"ISC/Websites"
	"embed"
	"encoding/json"
	"math/rand"
	"time"
)

//go:embed database.json
var f embed.FS

func readJSON() string {
	data, _ := f.ReadFile("database.json")
	return string(data)
}

type JSONWebsite struct {
	Domain string
	Rank   int
}

type DatabaseStructure struct {
	Websites []JSONWebsite
}

func GetWebsites() []JSONWebsite {
	database := readJSON()
	var website []JSONWebsite
	json.Unmarshal([]byte(database), &website)
	return website
}

func getRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func GetRandomWebsite() Websites.Website {
	websites := GetWebsites()
	length := len(websites)
	maxIndex := length - 1
	randomNumber := getRandomNumber(0, maxIndex)
	randomSite := websites[randomNumber]
	return Websites.Website{Domain: randomSite.Domain}
}
