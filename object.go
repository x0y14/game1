package main

import (
	"image/color"
)

type Object struct {
	Component []*Polygon
	Center    Position
}

func (o *Object) SetPosition(p Position) {
	o.Center = p
}

type Polygon struct {
	Figure
	Center Position
	Width  float32
	Height float32
	Radius float32
	Color  color.Color
}
type Figure int

const (
	Rectangle Figure = iota
	Circle
)

type Position struct {
	X float32
	Y float32
}
