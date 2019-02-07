package ports

import (
	"github.com/mao-wfs/mao-wfs/domain/model"
)

// VDIFPort represents the port of VDIF (VLBI Data Interchange Format).
type VDIFPort interface {
	CalcDeformation(vdif []byte) (*model.Deformation, error)
}
