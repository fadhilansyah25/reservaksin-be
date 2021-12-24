package citizen
import(
	"context"
)
type Domain struct{
	ID int
}

type Service interface {
	CreateToken(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, data *Domain) error
	GetById(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data Domain) error
}