package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record (row, entry) not found")
)

type Models struct {
	Movies MovieModel
}

type Directors struct {
	Movies MovieModel
}

func NewModelsDirectors(db *sql.DB) Directors {
	return Directors{
		Movies: MovieModel{DB: db},
	}
}

type Director struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Surname string   `json:"surname"`
	Awards  []string `json: "awards"`
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
