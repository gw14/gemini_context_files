package main

import (
	"math"
	"fmt"
)

//defining recangle
type Rectangle struct{
	Height float64
	Width float64
}

func (r *Rectangle) Area() float64{
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64{
	return (r.Height + r.Width) * 2
}

//defining Circle
type Circle struct{
	Radius float64
}

func (c Circle) Area() float64{
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float64{
	return 2 * c.Radius * math.Pi
}

//adding shape interface
type Shape interface{
	Area() float64
}

func PrintArea(s Shape) {
	fmt.Printf("The Area is: %f\n", s.Area())
}

func main(){

	c := Circle{
		Radius: 5,
	}
	r := Rectangle{
		Height: 4,
		Width: 3,
	}

	fmt.Printf("Circle: Area %f Perimeter %f\n",c.Area(),c.Perimeter())
	fmt.Printf("Recangle: Area %f Perimeter %f\n",r.Area(),r.Perimeter())
	PrintArea(&r)
	PrintArea(c)
	
}
