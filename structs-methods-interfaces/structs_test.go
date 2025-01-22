package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10, 10}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("Want '%.2f', got '%.2f'", want, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12.0, height: 6.0}, hasArea: 72},
		{name: "Circle", shape: Circle{radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12.0, height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got '%g', want '%g'", tt.shape, got, tt.hasArea)
			}
		})
	}

}
