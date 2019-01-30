package repository

import (
	"github.com/mao-wfs/maoctrl/domain/model"
)

// DeformationsRepository is the domain repository for deformation of the antenna.
type DeformationsRepository interface {
	Store(deformations *model.Deformations) error
	FindList() ([]*model.Deformations, error)
}
