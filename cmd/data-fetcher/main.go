package main

import (
	"fmt"
	"log"
	"os"

	"github.com/przant/zipcodes-api/utils"
)

const (
	INSERT = "USE us_zipcodes;\nINSERT INTO zipcodes(state, state_abbr, zipcode, county, city) VALUES"
)

func main() {
	zips, err := utils.FetchData()
	if err != nil {
		log.Fatalf("while feching the zipcodes data: %s", err)
	}

	file, err := os.OpenFile("insert.sql", os.O_RDWR|os.O_CREATE, 0650)
	if err != nil {
		log.Fatalf("while vcreating the file %q: %s", "insert.sql", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%s\n", INSERT)
	for _, v := range zips[1 : len(zips)-1] {
		fmt.Fprintf(file, "(%q,%q,%q,%q,%q),\n", v.State, v.StateAbbr, v.Zipcode, v.County, v.City)
	}
	v := zips[len(zips)-1]
	fmt.Fprintf(file, "(%q,%q,%q,%q,%q);\n", v.State, v.StateAbbr, v.Zipcode, v.County, v.City)
}
