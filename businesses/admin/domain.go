package admin

import (
	"time"
)

type Domain struct {
	Id        string
	Role      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Register(dataAdmin *Domain) (Domain, error)
	Login(username, password string) (string, error)
	GetByID(id string) (Domain, error)
}

type Repository interface {
	Register(dataAdmin *Domain) (Domain, error)
	GetByUsername(username string) (Domain, error)
	GetByID(id string) (Domain, error)
}
