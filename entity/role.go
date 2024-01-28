package entity

type Role struct {
	ID          int64  `json:"id"`
	RoleName    string `json:"roleName"`
	Users       []User
	Permissions []*Permission `gorm:"many2many:role_permissions;"` // Define many-to-many relationship
}
