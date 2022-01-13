package vaccine

import "time"

type Domain struct {
	Id         string
	NamaVaksin string
	Stok       int
	AdminID    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	Create(data *Domain) (Domain, error)
	Update(id string, data *Domain) (Domain, error)
	Delete(id string) (string, error)
	GetByID(id string) (Domain, error)
	FetchAll() ([]Domain, error)
	GetByAdminID(adminID string) ([]Domain, error)
}

type Repository interface {
	Create(data *Domain) (Domain, error)
	Update(id string, data *Domain) (Domain, error)
	Delete(id string) (string, error)
	GetByID(id string) (Domain, error)
	FetchAll() ([]Domain, error)
	GetByAdminID(adminID string) ([]Domain, error)
}
