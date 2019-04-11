package model

import (
	"time"
)

// Deformation represents deformation of the position on the antenna.
type Deformation struct {
	Time       time.Time
	PositionID PositionID
	Volume     float32
}

// NewDeformation returns a new deformation.
func NewDeformation(t time.Time, posID PositionID, v float32) *Deformation {
	return &Deformation{
		Time:       t,
		PositionID: posID,
		Volume:     v,
	}
}

// GetTime returns the time of the deformation.
func (d *Deformation) GetTime() time.Time { return d.Time }

// GetPositionID returns the position ID of the deformation.
func (d *Deformation) GetPositionID() PositionID { return d.PositionID }

// GetVolume returns the deformation volume.
func (d *Deformation) GetVolume() float32 { return d.Volume }
