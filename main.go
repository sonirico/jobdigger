package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"jobdigger/digger/empleoPonferradaOrg"
	"jobdigger/offer"
	"log"
)

func main() {
	empleoPonferradaOrgDigger := empleoPonferradaOrg.New("https://empleo.ponferrada.org/rss")
	result, err := empleoPonferradaOrgDigger.GetOffers()

	if err != nil {
		fmt.Println("Fetch error")
	}

	db, err := sql.Open("sqlite3", "./database/jobdigger.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for index, o := range result.Offers {
		fmt.Println(index, o)
		offer.NewSQliteOffer(&o, db).Insert()
	}

	fmt.Println(fmt.Sprintf("OK: %d \nFAILED: %d\nTOTAL: %d", result.Ok, result.Failed, result.Total()))
}
