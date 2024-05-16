package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Shape interface {
	SetArea(int)
	GetArea() int
}

type Circle struct {
	a int
}

func (c *Circle) SetArea(a int) {
	c.a = a
}

func (c Circle) GetArea() int {
	return c.a
}

func NewCircle(a int) Shape {
	return &Circle{a: a}
}

type Rectangle struct {
	a int
}

func (c *Rectangle) SetArea(a int) {
	c.a = a
}

func (c Rectangle) GetArea() int {
	return c.a
}

type NTC struct {
	C Shape
}

func NewNTC(c Shape) *NTC {
	return &NTC{
		C: c,
	}
}

func TestSingleton(t *testing.T) {
	area := 5

	// register the instance
	Register[Shape](NewCircle(area))
	// override
	Register(
		func() Shape {
			return &Rectangle{a: area}
		},
	)

	// register with the arg as abstraction already have in container.
	Register(
		func(c Shape) *NTC {
			return NewNTC(c)
		},
	)

	UseInOtherPlace(t, area)

	var myRectangle Shape
	myRectangle = ResolverMust[Shape]()
	a := myRectangle.GetArea()
	assert.Equal(t, area, a)

	myNTC := ResolverMust[*NTC]()
	println(myNTC.C.GetArea())
}

func UseInOtherPlace(t *testing.T, area int) {
	// myCircle, err := Resolver[Shape]()
	// if err != nil {
	// 	panic(err)
	// }
	myCircle := ResolverMust[Shape]()
	a := myCircle.GetArea()
	println(a)
	assert.Equal(t, area, a)
}
