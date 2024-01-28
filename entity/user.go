package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                      int64      `json:"-"`    // Primary key, excluding from JSON
	UUID                    uuid.UUID  `json:"uuid"` // UUID for the user
	Phone                   string     `json:"phone"`
	Email                   string     `json:"email"`
	Password                string     `json:"-"` // Excluding password from JSON
	TotpSecret              string     `json:"-"` // Excluding TOTP secret from JSON
	IsUsingMfa              bool       `json:"isUsingMfa"`
	IsLoggedIn              bool       `json:"isLoggedIn"`
	IsActive                bool       `json:"isActive"`
	IsEnabled               bool       `json:"isEnabled"`
	IsNotExpired            bool       `json:"isNotExpired"`
	IsNotLocked             bool       `json:"isNotLocked"`
	IsCredentialsNonExpired bool       `json:"isCredentialsNonExpired"`
	IsDeleted               bool       `json:"isDeleted"`
	JoinDate                time.Time  `json:"joinDate"`
	LastLoginDate           *time.Time `json:"lastLoginDate,omitempty"`
	LastUpdated             *time.Time `json:"lastUpdated,omitempty"`
	DeletedAt               *time.Time `json:"deletedAt,omitempty"`
	Roles                   []Role     `gorm:"many2many:user_roles;"` // Using GORM for ORM
}
