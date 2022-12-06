package geometry

type Line struct {
	Origin Point
	Terminus Point
}

// determine whether the specified line is horizontal
func (l Line) Horizontal() bool {
	return l.Origin.Y == l.Terminus.Y
}

// determine whether the specified line is vertical
func (l Line) Vertical() bool {
	return l.Origin.X == l.Terminus.X
}

func (l Line) Intersects(j Line) (intersects bool) {
	if ((l.Horizontal() && j.Horizontal()) || (l.Vertical() && j.Vertical())) {
		return false
	}

	if (l.Origin.X == l.Terminus.X) {
			return !(j.Origin.X == j.Terminus.X && l.Origin.X != j.Origin.X)
	} else if (j.Origin.X == j.Terminus.X) {
			return true
	} else {
		// neither line is parallel to the y-axis
		m1 := (l.Origin.Y - l.Terminus.Y) / (l.Origin.X - l.Terminus.X);
		m2 := (j.Origin.Y - j.Terminus.Y) / (l.Origin.X - j.Terminus.X);
		return m1 != m2
	}
}

func (l Line) Slope() float64 {
  return (l.Terminus.Y - l.Origin.Y) / (l.Terminus.X - l.Origin.X)
}

// // fill in an array of points on a provided line
// func (l Line) FillInPoints() []Point {
// 	var points []Point

// 	if l.Horizontal() {
// 		min := math.Min(float64(l.Origin.X), float64(l.Terminus.X))
// 		max := math.Max(float64(l.Origin.X), float64(l.Terminus.X))
// 		for i := min; i <= max; i++ {
// 			points = append(points, Point{i, l.Origin.Y})
// 		}
// 	} else if l.Vertical() {
// 		min := math.Min(float64(l.Origin.Y), float64(l.Terminus.Y))
// 		max := math.Max(float64(l.Origin.Y), float64(l.Terminus.Y))
// 		for i := min; i <= max; i++ {
// 			points = append(points, Point{l.Origin.X, i})
// 		}
// 	} else {

// 		points = append(points, Point{1,1})
// 		if l.Origin.X > l.Terminus.X {
// 			xstep := -1
// 		} else {
// 			xstep := 1
// 		}
// 		if l.Origin.Y > l.Terminus.Y {
// 			ystep := -1
// 		} else {
// 			ystep := 1
// 		}
// 		xpoint := l.Origin.X
// 		ypoint := l.Origin.Y
// 		// while xpoint != l.Terminus.X
// 		// 	points << Point.new(xpoint, ypoint)
// 		// 	xpoint += xstep
// 		// 	ypoint += ystep
// 		// end
// 		points << Point{l.Terminus.X, l.Terminus.Y}
// 	}

// 	return points
// }
