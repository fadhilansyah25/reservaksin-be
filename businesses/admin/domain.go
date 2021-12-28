package admin

import (
	"time"
)

type Domain struct {
	Id        int
	Role      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
