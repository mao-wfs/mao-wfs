package model

import (
	"time"
)

// Deformation represents a deformation of the antenna.
type Deformation struct {
	Time       time.Time
	PositionID int64
	Volume     float32
}
