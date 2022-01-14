package formas

import (
	"math"
	"testing"
)

func TestGetArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{Width: 10, Height: 5}
		area := GetArea(r)
		if area != 50 {
			t.Errorf("Esperado 50, mas obteve %f", area)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{Radius: 5}
		area := GetArea(c)
		if math.Round(area*10)/10 != 78.5 {
			t.Errorf("Esperado 78.5, mas obteve %f", area)
		}
	})
}
