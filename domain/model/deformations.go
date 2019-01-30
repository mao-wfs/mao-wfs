package model

import (
	"time"
)

// Deformation represents a deformation of the antenna.
type Deformation struct {
	PositionID int64
	Volume     float32
}

// Deformations represents the total deformation of the antenna.
type Deformations struct {
	Time        time.Time
	Deformation []*Deformation
}
