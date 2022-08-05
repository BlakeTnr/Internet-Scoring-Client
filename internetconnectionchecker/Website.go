package internetchecker

import (
	"fmt"
	"isc/database"
	"math/rand"
	"net/http"
	"time"
)

type Website struct {
	Domain string
}

func RandomWebsite() Website {
	rand.Seed(time.Now().UnixNano())
	numberOfWebsites := database.GetNumberOfWebsites()
	randomIndex := rand.Intn(numberOfWebsites)
	domain := database.GetDomainByIndex(randomIndex)
	return Website{domain}
}

func (website *Website) GetURL() string {
	url := "https://" + website.Domain + "/"
	return url
}

func (website *Website) VerifyConnection() bool {
	_, err := http.Get(website.GetURL())

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
