package models

import (
	"strings"
	"time"
)

const (
	RoleAdmin = "admin"
)

type UserDTO struct {
	ID           uint      `json:"id"`
	AuthID       string    `json:"authID,omitempty"`
	AuthProvider string    `json:"authProvider,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	Email        string    `json:"email,omitempty"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Scopes       string    `json:"scopes"`
}

type User struct {
	Model
	AuthID       string  `json:"authID,omitempty"`
	AuthProvider string  `json:"authProvider,omitempty"`
	FirstName    string  `json:"firstName" validate:"min=1,max=64"`
	LastName     string  `json:"lastName" validate:"min=2,max=64"`
	Email        string  `json:"email,omitempty" gorm:"uniqueIndex,size:255" validate:"email,max=64"` // todo: required if no mobile
	History      []Login `json:"history,omitempty" gorm:"foreignKey:CreatedByID"`
	Password     string  `json:"password,omitempty"`
	Scopes       string  `json:"scopes"`
}

func (u *User) Is(role string) bool {
	return strings.Contains(u.Scopes, role)
}

func (u *User) IsBasic() bool {
	return !strings.Contains(u.Scopes, RoleAdmin)
}