package handler

import (
	"apikeys/internal/app/relationtype/service"
	"database/sql"
	"encoding/json"
	"net/http"
)

const nameParam = "name"

type RelationTypeHandler struct {
	DB *sql.DB
}

func (h *RelationTypeHandler) GetRelationTypeByName(w http.ResponseWriter, r *http.Request) {
	service := &service.RelationService{DB: h.DB}

	name := r.URL.Query().Get(nameParam)

	RelationType, err := service.GetRelationByName(name)

	if RelationType == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(RelationType)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(responseJSON)
}
