package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("Rectangle: perimeter", func(t *testing.T) {
		r := rectangle{10.0, 15.0}
		got := r.perimeter()
		want := 50.0

		if got != want {
			t.Errorf("Got -->%.2f<--, want -->%.2f<--", got, want)
		}
	})

	t.Run("Rectangle: area", func(t *testing.T) {
		r := rectangle{10.0, 15.0}
		got := r.area()
		want := 150.0

		if got != want {
			t.Errorf("Got -->%.2f<--, want -->%.2f<--", got, want)
		}
	})

	t.Run("Circle: circumference", func(t *testing.T) {
		c := circle{10.0}
		got := c.perimeter()
		want := 62.83185307179586476925286766559

		if got != want {
			t.Errorf("Got -->%.2f<--, want -->%.2f<--", got, want)
		}
	})

	t.Run("Circle: area", func(t *testing.T) {
		c := circle{10.0}
		got := c.area()
		want := 314.15926535897932384626433832795

		if got != want {
			t.Errorf("Got -->%.2f<--, want -->%.2f<--", got, want)
		}
	})

	t.Run("Triangle: area", func(t *testing.T) {
		tri := triangle{15.0, 10.0}
		got := tri.area()
		want := 75.0

		if got != want {
			t.Errorf("Got -->%.2f<--, want -->%.2f<--", got, want)
		}
	})
}
