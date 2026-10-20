package shapes

import "testing"

func TestShapes(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want .2%f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
