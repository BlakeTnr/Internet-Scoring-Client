package Checks

import (
	"ISC/Database"
	"ISC/Websites"
	"net/http"
)

func RunCheck() []Result {
	var results []Result
	for i := 0; i < 3; i++ {
		website := Database.GetRandomWebsite()
		result := websiteCheck(website)
		results = append(results, result)
	}
	return results
}

func websiteCheck(website Websites.Website) Result {
	url := domainToUrl(website.Domain)
	response, err := http.Get(url)
	if err != nil {
		return Result{false, "Couldn't reach site"}
	}
	defer response.Body.Close()
	return Result{true, ""}
}

func domainToUrl(domain string) string {
	url := "https://" + domain + "/"
	return url
}
