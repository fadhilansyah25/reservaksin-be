package news

import (
	"time"
)

type Domain struct {
	Id        int
	Image     string
	Title     string
	Url       string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
