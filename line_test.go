package geometry

import (
	"math"
	"testing"
)

func TestVertical(t *testing.T){
	p1 := Point{0,0}
	p2 := Point{0,5}

	l := Line{p1, p2}

	if got := l.Vertical(); got != true {
		t.Errorf("Line is vertical, got %v\n", got)
	}
}

func TestHorizontal(t *testing.T) {
	p1 := Point{0,0}
	p2 := Point{5,0}

	l := Line{p1,p2}

	if got := l.Horizontal(); got != true {
		t.Errorf("Line is horizontal, got %v\n", got)
	}
}

func TestIntersects(t *testing.T) {
	l1 := Line {
		Point {0,3},
		Point {3,0},
	}
	l2 := Line{
		Point {3,3},
		Point {0,0},
	}

	if got := l1.Intersects(l2); got != true {
		t.Errorf("Lines intersect, got %v\n", got)
	}
}

func TestSlope(t *testing.T) {
	l1 := Line {
		Point {0,3},
		Point {3,0},
	}
	if got := l1.Slope(); got != -1 {
		t.Errorf("got %v\n", got)
	}

	l2 := Line {
		Point {0,0},
		Point {3,3},
	}
	if got := l2.Slope(); got != 1 {
		t.Errorf("got %v\n", got)
	}

	// horizontal
	l3 := Line {
		Point {0,0},
		Point {3,0},
	}
	if got := l3.Slope(); got != 0 {
		t.Errorf("got %v\n", got)
	}

	// vertical
	l4 := Line {
		Point {0,0},
		Point {0,3},
	}
	if got := l4.Slope(); got != math.Inf(1) {
		t.Errorf("got %v\n", got)
	}
}

// func TestFillInPoints(t *testing.T) {
// 	// horizontal
// 	l1 := Line {
// 		Point {0,0},
// 		Point {3,0},
// 	}
// 	l1_expected := [4]Point{{0,0},{1,0},{2,0},{3,0}}
// 	if l1_got := l1.FillInPoints(); reflect.DeepEqual(l1_got, l1_expected) == false {
// 		t.Errorf("got %v, expected %v\n", l1_got, l1_expected)
// 	}
// 	// vertical
// 	l2 := Line {
// 		Point {0,0},
// 		Point {0,3},
// 	}
// 	l2_expected := [4]Point{{0,0},{0,1},{0,2},{0,3}}
// 	if l2_got := l2.FillInPoints(); reflect.DeepEqual(l2_got, l2_expected) == false {
// 		t.Errorf("got %v, expected %v\n", l2_got, l2_expected)
// 	}
// 	// neither
// 	l3 := Line {
// 		Point {0,0},
// 		Point {3,3},
// 	}
// 	l3_expected := [4]Point{{0,0},{1,1},{2,2},{3,3}}
// 	if l3_got := l3.FillInPoints(); reflect.DeepEqual(l3_got, l3_expected) == false {
// 		t.Errorf("got %v, expected %v\n", l3_got, l3_expected)
// 	}
// }
