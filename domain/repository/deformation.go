package repository

import (
	"time"

	"github.com/mao-wfs/maoctrl/domain/model"
)

// DeformationRepository is the domain repository for deformation of the antenna.
type DeformationRepository interface {
	Store(deformations *model.Deformation) error
	FindByPositionID(posID int64, beginTime, endTime time.Time) ([]*model.Deformation, error)
}