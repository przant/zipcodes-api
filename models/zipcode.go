package models

type Zipcode struct {
	StateFIPS string `json:"state_fips"`
	State     string `json:"state"`
	StateAbbr string `json:"state_abbr"`
	Zipcode   string `json:"zipcode"`
	County    string `json:"county"`
	City      string `json:"city"`
}
