package primitives

import (
	"fmt"
	"math"
)

// Round will round all the coordinates inplace of the geometry to the given factor.
// The default is 6 decimal places.
func Round(g Geometry, factor ...int) Geometry {
	if g == nil {
		return nil
	}

	f := float64(DefaultRoundingFactor)
	if len(factor) > 0 {
		f = float64(factor[0])
	}

	switch g := g.(type) {
	case Point:
		roundPoint(g, f)
		return g
	case MultiPoint:
		roundPoints([]Point(g), f)
		return g
	case LineString:
		roundPoints([]Point(g), f)
		return g
	case MultiLineString:
		for _, ls := range g {
			roundPoints([]Point(ls), f)
		}
		return g
	case Ring:
		roundPoints([]Point(g), f)
		return g
	case Polygon:
		for _, r := range g {
			roundPoints([]Point(r), f)
		}
		return g
	case MultiPolygon:
		for _, p := range g {
			for _, r := range p {
				roundPoints([]Point(r), f)
			}
		}
		return g
	case Collection:
		for i := range g {
			g[i] = Round(g[i], int(f))
		}
		return g
	case Bound:
		g.Min[0] = math.Round(g.Min[0]*f) / f
		g.Min[1] = math.Round(g.Min[1]*f) / f
		g.Max[0] = math.Round(g.Max[0]*f) / f
		g.Max[1] = math.Round(g.Max[1]*f) / f
		return g
	}

	panic(fmt.Sprintf("geometry type not supported: %T", g))
}

func roundPoints(ps []Point, f float64) {
	for i := range ps {
		roundPoint(ps[i], f)
	}
}

func roundPoint(p Point, f float64) {
	for i := range p {
		p[i] = math.Round(p[i]*f) / f
	}
}
