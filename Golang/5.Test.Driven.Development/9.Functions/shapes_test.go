package shapes

import "testing"

// TestPerimeter of the rectangle
func TestPerimeter(t *testing.T) {
	got := perimeter(10.0, 15.0)
	want := 50.0

	if got != want {
		t.Errorf("Got -->%.2f<--, want -->%.2f<--", got, want)
	}
}

// TestArea of the rectangle
func TestArea(t *testing.T) {
	got := area(10.0, 15.0)
	want := 150.0

	if got != want {
		t.Errorf("Got -->%.2f<--, want -->%.2f<--", got, want)
	}
}
