package service

import (
	m "apikeys/internal/app/relationtype/model"
	"database/sql"
	"fmt"
	"strings"
)

type RelationService struct {
	DB *sql.DB
}

func (h *RelationService) GetRelationByName(name string) (*m.RelationType, error) {
	nameCanonical := strings.ToLower(strings.TrimSpace(name))

	var relationType m.RelationType

	err := h.DB.QueryRow(`
						SELECT 
						    id, name, name_canonical, comment
						FROM 
						    relation_type 
						WHERE 
						    name_canonical = $1`,
		nameCanonical).Scan(&relationType.ID, &relationType.Name, &relationType.NameCanonical, &relationType.Comment)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("RelationType with name %s not found", name)
	}

	if err != nil {
		return nil, err
	}

	return &relationType, nil

}
