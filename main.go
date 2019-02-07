package main

import (
	"database/sql"
	"fmt"
	"jobdigger/digger/empleoPonferradaOrg"
	offer2 "jobdigger/offer"
	_ "github.com/mattn/go-sqlite3"
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

	for index, offer := range result.Offers {
		fmt.Println(index, offer)
		insertStmt := offer2.NewSQliteOffer(&offer).Insert()

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := tx.Prepare(insertStmt)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec()
		tx.Commit()
	}


	fmt.Println(fmt.Sprintf("OK: %d \nFAILED: %d\nTOTAL: %d", result.Ok, result.Failed, result.Total()))
}
