package models

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
