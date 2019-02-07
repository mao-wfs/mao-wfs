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

// GetTime gets the time of the Deformation.
func (d *Deformation) GetTime() time.Time { return d.Time }

// GetPositionID gets the position ID of the Deformation.
func (d *Deformation) GetPositionID() int64 { return d.PositionID }

// GetVolume gets the volume of the Deformation.
func (d *Deformation) GetVolume() float32 { return d.Volume }
