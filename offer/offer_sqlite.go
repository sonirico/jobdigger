package offer

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	DbTable = "offer"
)

type SQliteOffer struct {
	db    *sql.DB
	Offer *Offer
}

func NewSQliteOffer(offer *Offer, db *sql.DB) *SQliteOffer {
	return &SQliteOffer{Offer: offer, db: db}
}

func (so *SQliteOffer) Insert() {
	insertStmt := fmt.Sprintf(`
		INSERT INTO %s
		(title, permalink, description, published_at)
		VALUES (?, ?, ?, ?)
	`, DbTable)
	tx, err := so.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(insertStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(so.Offer.Title, so.Offer.Link, so.Offer.Description, so.Offer.PublishedAtUnix())
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
