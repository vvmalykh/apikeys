package model

type APIKeyRelation struct {
	APIKeyID     string `db:"api_key_id"`
	RelationID   string `db:"relation_id"`
	RelationType string `db:"relation_type_id"`
	RelatedID    string `db:"related_id"`
}
