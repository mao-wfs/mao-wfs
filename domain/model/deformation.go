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

// GetTime returns "Time".
func (d *Deformation) GetTime() time.Time { return d.Time }

// GetPositionID returns "PositionID".
func (d *Deformation) GetPositionID() PositionID { return d.PositionID }

// GetVolume returns "Volume".
func (d *Deformation) GetVolume() float32 { return d.Volume }
