package models

type Statement struct {
	JsonTransaction []JsonTransaction `json:"transactions"`
}

// test commit 6