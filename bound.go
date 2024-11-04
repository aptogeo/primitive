package primitives

import (
	"math"
)

var emptyBound = Bound{Min: [2]float64{1, 1}, Max: [2]float64{-1, -1}}

// A Bound represents a closed box or rectangle.
// To create a bound with two points you can do something like:
//
//	orb.MultiPoint{p1, p2}.Bound()
type Bound struct {
	Min, Max [2]float64
}

// GeoJSONType returns the GeoJSON type for the object.
func (b Bound) GeoJSONType() string {
	return "Polygon"
}

// Dimensions returns 2 because a Bound is a 2d object.
func (b Bound) Dimensions() int {
	return 2
}

// ToPolygon converts the bound into a Polygon object.
func (b Bound) ToPolygon() Polygon {
	return Polygon{b.ToRing()}
}

// ToRing converts the bound into a loop defined
// by the boundary of the box.
func (b Bound) ToRing() Ring {
	return Ring{
		Point{b.Min[0], b.Min[1]},
		Point{b.Max[0], b.Min[1]},
		Point{b.Max[0], b.Max[1]},
		Point{b.Min[0], b.Max[1]},
		Point{b.Min[0], b.Min[1]},
	}
}

// Extend grows the bound to include the new point.
func (b Bound) Extend(point Point) Bound {
	// already included, no big deal
	if b.Contains(point) {
		return b
	}

	return Bound{
		Min: [2]float64{
			math.Min(b.Min[0], point[0]),
			math.Min(b.Min[1], point[1]),
		},
		Max: [2]float64{
			math.Max(b.Max[0], point[0]),
			math.Max(b.Max[1], point[1]),
		},
	}
}

// Union extends this bound to contain the union of this and the given bound.
func (b Bound) Union(other Bound) Bound {
	if other.IsEmpty() {
		return b
	}

	b = b.Extend(Point{other.Min[0], other.Min[1]})
	b = b.Extend(Point{other.Max[0], other.Max[1]})
	b = b.Extend(other.LeftTop())
	b = b.Extend(other.RightBottom())

	return b
}

// Contains determines if the point is within the bound.
// Points on the boundary are considered within.
func (b Bound) Contains(point Point) bool {
	if point[1] < b.Min[1] || b.Max[1] < point[1] {
		return false
	}

	if point[0] < b.Min[0] || b.Max[0] < point[0] {
		return false
	}

	return true
}

// Intersects determines if two bounds intersect.
// Returns true if they are touching.
func (b Bound) Intersects(bound Bound) bool {
	if (b.Max[0] < bound.Min[0]) ||
		(b.Min[0] > bound.Max[0]) ||
		(b.Max[1] < bound.Min[1]) ||
		(b.Min[1] > bound.Max[1]) {
		return false
	}

	return true
}

// Pad extends the bound in all directions by the given value.
func (b Bound) Pad(d float64) Bound {
	b.Min[0] -= d
	b.Min[1] -= d

	b.Max[0] += d
	b.Max[1] += d

	return b
}

// Center returns the center of the bounds by "averaging" the x and y coords.
func (b Bound) Center() Point {
	return Point{
		(b.Min[0] + b.Max[0]) / 2.0,
		(b.Min[1] + b.Max[1]) / 2.0,
	}
}

// Top returns the top of the bound.
func (b Bound) Top() float64 {
	return b.Max[1]
}

// Bottom returns the bottom of the bound.
func (b Bound) Bottom() float64 {
	return b.Min[1]
}

// Right returns the right of the bound.
func (b Bound) Right() float64 {
	return b.Max[0]
}

// Left returns the left of the bound.
func (b Bound) Left() float64 {
	return b.Min[0]
}

// LeftTop returns the upper left point of the bound.
func (b Bound) LeftTop() Point {
	return Point{b.Left(), b.Top()}
}

// RightBottom return the lower right point of the bound.
func (b Bound) RightBottom() Point {
	return Point{b.Right(), b.Bottom()}
}

// IsEmpty returns true if it contains zero area or if
// it's in some malformed negative state where the left point is larger than the right.
// This can be caused by padding too much negative.
func (b Bound) IsEmpty() bool {
	return b.Min[0] > b.Max[0] || b.Min[1] > b.Max[1]
}

// IsZero return true if the bound just includes just null island.
func (b Bound) IsZero() bool {
	return b.Max == [2]float64{0, 0} && b.Min == [2]float64{0, 0}
}

// Bound returns the the same bound.
func (b Bound) Bound() Bound {
	return b
}

// Equal returns if two bounds are equal.
func (b Bound) Equal(c Bound) bool {
	return b.Min == c.Min && b.Max == c.Max
}
