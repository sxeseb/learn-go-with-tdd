package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := Perimeter(r)
		want := 40.0
		assertNumbers(t, want, got)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v: want %g, got %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}

func assertNumbers(t testing.TB, want, got float64) {
	t.Helper()
	if want != got {
		t.Errorf("want %g got %g", want, got)
	}
}
