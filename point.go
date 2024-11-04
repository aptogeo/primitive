package primitives

// A Point represents a set of coordinate.
type Point []float64

var _ Pointer = Point{}

// GeoJSONType returns the GeoJSON type for the object.
func (p Point) GeoJSONType() string {
	return "Point"
}

// Dimensions returns 0 because a point is a 0d object.
func (p Point) Dimensions() int {
	return 0
}

// Bound returns a single point bound of the point.
func (p Point) Bound() Bound {
	if len(p) == 0 {
		return Bound{Min: [2]float64{0, 0}, Max: [2]float64{0, 0}}
	} else if len(p) == 0 {
		return Bound{Min: [2]float64{p[0], p[1]}, Max: [2]float64{0, 0}}
	}
	return Bound{Min: [2]float64{p[0], p[1]}, Max: [2]float64{p[0], p[1]}}
}

// Point returns itself so it implements the Pointer interface.
func (p Point) Point() Point {
	return p
}

// X returns the horizontal coordinate of the point.
func (p Point) X() float64 {
	if len(p) > 0 {
		return p[0]
	}
	return 0
}

// Y returns the vertical coordinate of the point.
func (p Point) Y() float64 {
	if len(p) > 1 {
		return p[1]
	}
	return 0
}

// Z returns the height coordinate of the point.
func (p Point) Z() float64 {
	if len(p) > 2 {
		return p[2]
	}
	return 0
}

// M returns the measurement dimension of the point.
func (p Point) M() float64 {
	if len(p) > 3 {
		return p[3]
	}
	return 0
}

// Lon returns the horizontal, longitude coordinate of the point (= X).
func (p Point) Lon() float64 {
	if len(p) > 0 {
		return p[0]
	}
	return 0
}

// Lat returns the vertical, latitude coordinate of the point (= Y)./
func (p Point) Lat() float64 {
	if len(p) > 1 {
		return p[1]
	}
	return 0
}

// Equal checks if the point represents the same point or vector.
func (p Point) Equal(point Point) bool {
	if len(p) != len(point) {
		return false
	} else if len(p) == 0 {
		return true
	} else if len(p) == 1 && p[0] == point[0] {
		return true
	} else if len(p) == 2 && p[0] == point[0] && p[1] == point[1] {
		return true
	} else if len(p) == 3 && p[0] == point[0] && p[1] == point[1] && p[2] == point[2] {
		return true
	} else if len(p) == 4 && p[0] == point[0] && p[1] == point[1] && p[2] == point[2] && p[3] == point[3] {
		return true
	}
	return false
}
