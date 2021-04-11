package models

type Value struct {
	CreatedBy string `db:"created_by"`
	Key       string `db:"key"`
	Value     string `db:"value"`
	Roles     string `db:"roles"`
}
