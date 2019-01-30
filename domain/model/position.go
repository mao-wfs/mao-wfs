package model

// Position represents a position on the antenna.
type Position struct {
	ID          int64
	Coordinates Coordinates
}

// Coordinates represents the coordinates of the antenna.
type Coordinates struct {
	Radius float32
	Angle  float32
}
