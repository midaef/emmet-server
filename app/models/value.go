package models

type Value struct {
	ID           uint64   `json:"id"            db:"id"`
	Key          string   `json:"key"           db:"key"`
	Value        string   `json:"value"         db:"value"`
	CreatedBy    uint64   `json:"created_by"    db:"created_by"`
	Config       uint64   `json:"config"        db:"config"`
	AllowedRoles []uint64 `json:"allowed_roles" db:"allowed_roles"`
}
