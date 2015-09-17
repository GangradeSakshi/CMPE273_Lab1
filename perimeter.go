package main

import ("fmt"; "math")

type Circle struct {
  r float64
}

func (c *Circle) perimeter() float64 {
  return math.Pi * c.r*2
}

type Rectangle struct {
  x,y float64
}
func (r *Rectangle) perimeter() float64 {
  l := r.x + r.y
  return l * 2
}

type Shape interface {
  perimeter() float64
}

func main() {
  c := Circle{5}
  fmt.Println(c.perimeter())
  r := Rectangle{2,8}
  fmt.Println(r.perimeter())
}
