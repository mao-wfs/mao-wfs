package model

// Position represents a position on the antenna.
type Position struct {
	Coordinates Coordinates
	ID          PositionID
}

// PositionID is the ID of a position on the antenna.
type PositionID int

// Coordinates represents the coordinates of the antenna.
type Coordinates struct {
	Radius float32
	Angle  float32
}
