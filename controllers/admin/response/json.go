package response

import (
	"ca-reservaksin/businesses/admin"
	"time"
)

type Admin struct {
	Id        string    `json:"id"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain admin.Domain) *Admin {
	return &Admin{
		Id:        domain.Id,
		Role:      domain.Role,
		Username:  domain.Username,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
