package models

type Role struct {
	CreatedBy   string `db:"created_by"`
	CreateUser  bool `db:"create_user"`
	CreateRole  bool `db:"create_role"`
	CreateValue bool `db:"create_value"`
	DeleteUser  bool `db:"delete_user"`
	DeleteRole  bool `db:"delete_role"`
	DeleteValue bool `db:"delete_value"`
	Role        string `db:"user_role"`
}

type Permissions struct {
	CreateUser  bool `db:"create_user"`
	CreateRole  bool `db:"create_role"`
	CreateValue bool `db:"create_value"`
	DeleteUser  bool `db:"delete_user"`
	DeleteRole  bool `db:"delete_role"`
	DeleteValue bool `db:"delete_value"`
}