package dto

import (
	"Gin-Auth/entity"
	"time"
)

type SignUp struct {
	Phone      string
	Email      string
	Password   string
	JoinDate   time.Time
	IsUsingMfa bool
	Role       entity.Role
	Permission entity.Permission
}
