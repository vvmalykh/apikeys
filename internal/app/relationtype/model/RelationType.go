package model

import "database/sql"

type RelationType struct {
	ID            string         `json:"id" db:"id"`
	Name          string         `json:"name" db:"name"`
	NameCanonical string         `json:"-" db:"name_canonical"`
	Comment       sql.NullString `json:"comment" db:"comment"`
}
