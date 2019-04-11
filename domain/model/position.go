package model

// Position represents a position on the antenna.
type Position struct {
	ID          PositionID
	Coordinates *Coordinates
}

// PositionID is an ID of a position on the antenna.
type PositionID int

// GetID returns the position ID.
func (p *Position) GetID() PositionID { return p.ID }

// GetCoordinates returns the coordinates at the position.
func (p *Position) GetCoordinates() *Coordinates { return p.Coordinates }

// Coordinates represents the coordinates of the antenna.
type Coordinates struct {
	Radius float32
	Angle  float32
}

// GetRadius returns the radius at the position.
func (c *Coordinates) GetRadius() float32 { return c.Radius }

// GetAngle returns the angle at the position.
func (c *Coordinates) GetAngle() float32 { return c.Angle }
