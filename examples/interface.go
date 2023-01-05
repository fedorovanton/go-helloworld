package examples

import (
	"fmt"
	"math"
)

type Shape interface {
	// Композиция интерфейсов
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLength float32
}

func PrintShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

func (s Square) Perimeter() float32 {
	return s.sideLength * 4
}

func NewSquare(length float32) Square {
	return Square{
		sideLength: length,
	}
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func NewCircle(radius float32) Circle {
	return Circle{
		radius: radius,
	}
}

func PrintInterface(i interface{}) {
	// type switch
	switch value := i.(type) {
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("bool", value)
	default:
		fmt.Println("unknown type", value)
	}

	// Проверка интерфейса, что он типа string
	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not string")
		return
	} else {
		fmt.Println(str)
	}

	fmt.Printf("%+v\n", i)
}

func RunInterface() {
	// Работа с интерфейсами
	square := NewSquare(5)
	PrintShapeArea(square)
	PrintInterface(square)

	circle := NewCircle(8)
	PrintShapeArea(circle)
	PrintShapeArea(circle)
}
