package vaccine

import "time"

type Domain struct {
	Id         int
	NamaVaksin string
	Stok       int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	Create(data *Domain) (Domain, error)
	Update(id int, data *Domain) (Domain, error)
	Delete(id int) (string, error)
	GetByID(id int) (Domain, error)
	FetchAll() ([]Domain, error)
}

type Repository interface {
	Create(data *Domain) (Domain, error)
	Update(id int, data *Domain) (Domain, error)
	Delete(id int) (string, error)
	GetByID(id int) (Domain, error)
	FetchAll() ([]Domain, error)
}
