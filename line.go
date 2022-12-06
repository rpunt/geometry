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
