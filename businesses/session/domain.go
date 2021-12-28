package session

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type Domain struct {
	Id                int
	HealthFacilitesID int
	NameSession       string
	Kapasitas         int
	VaksinID          int
	Tanggal           string
	Tahap             string
	StartSession      timestamp.Timestamp
	EndSession        timestamp.Timestamp
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
