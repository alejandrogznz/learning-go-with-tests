package smi

import (
	"math"
)

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Width float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Perimeter() float64{
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return c.Radius*c.Radius*math.Pi
}

func (t Triangle) Area() float64 {
	return t.Width * t.Height / 2.0
}
