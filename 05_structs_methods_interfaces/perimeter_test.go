package structs_methods_interfaces


import (
  "testing"
  "math"
)


type Rectangle struct {
  Height float64
  Width float64
}


func (r Rectangle) Area() float64 {
  return r.Height * r.Width
}


type Circle struct {
  Radius float64
}


func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}


type Triangle struct {
  Base float64
  Height float64
}


func (t Triangle) Area() float64 {
  return 0.5 * (t.Base * t.Height)
}


type Shape interface {
  Area() float64
}


func TestPerimeter(t *testing.T) {
  rect := Rectangle{Height: 10.0, Width: 20.0}
  got := Perimeter(rect)
  want := 60.0
  if got != want {
    t.Errorf("want: %.2f, got: %.2f", want, got)
  }
}


func TestArea(t *testing.T) {
  areaTests := []struct {
    name string
    shape Shape
    hasArea float64
  }{
    {"Rectangle", Rectangle{Height: 10.0, Width: 20.0}, 200.0},
    {"Circle", Circle{Radius: 20}, 1256.6370614359173},
    {"Triangle", Triangle{Base: 12, Height: 6}, 36.0},
  }

  for _, test := range areaTests {
    t.Run(test.name, func(t *testing.T){
      got := test.shape.Area()
      if got != test.hasArea {
        t.Errorf("%#v want: %g, got: %g", test.shape, test.hasArea, got)
      }
    })
  }
}
