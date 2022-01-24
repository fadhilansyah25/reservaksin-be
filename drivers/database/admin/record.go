package admin

import (
	"ca-reservaksin/businesses/admin"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Id        string    `json:"id" gorm:"PrimaryKey; Not Null"`
	Role      string    `json:"role"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (rec *Admin) ToDomain() admin.Domain {
	return admin.Domain{
		Id:        rec.Id,
		Role:      rec.Role,
		Username:  rec.Username,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func FromDomain(domain admin.Domain) *Admin {
	return &Admin{
		Id:       domain.Id,
		Role:     domain.Role,
		Username: domain.Username,
		Password: domain.Password,
	}
}
