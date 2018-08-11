package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	assert := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		assert(t, got, want)
	})

}

func TestAera(t *testing.T) {

	assert := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.GetAera()
		want := 100.0

		assert(t, got, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.GetAera()
		want := 314.1592653589793

		assert(t, got, want)
	})

}
