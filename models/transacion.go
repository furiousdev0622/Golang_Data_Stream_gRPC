package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type JsonTransaction struct {
	ID          string     `json:"id"`
	Amount      string     `json:"amount"`
	Date        string     `json:"date"`
	Description string     `json:"description"`
	// Description string     `json:"description"`
}

// test commit 4

type Transaction struct {
	Identity          string     `gorm:"primary_key"`
	Amount      string     `gorm:"size:12"`
	Date        string     `gorm:"size:10"`
}
