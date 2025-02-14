package models

import (
	"math/rand"
	"time"
)

var companies = []struct {
	CompanyID   string
	CompanyName string
}{
	{"C01", "Uber"},
	{"C02", "PickMe"},
	{"C03", "REDBOX"},
}

func GetCompanyDetail(id string) (string, string, bool) {
	for _, company := range companies {
		if company.CompanyID == id {
			return company.CompanyID, company.CompanyName, true
		}
	}
	return "", "", false
}

func GetRandomCompany() (string, string) {
	rand.Seed(time.Now().UnixNano())

	// Select a random company
	randomIndex := rand.Intn(len(companies))
	selectedCompany := companies[randomIndex]

	return selectedCompany.CompanyID, selectedCompany.CompanyName
}
