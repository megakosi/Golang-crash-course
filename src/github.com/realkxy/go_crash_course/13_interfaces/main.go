package main
import (
	"fmt"
	"math"
)


type Shape interface {

	area() float64

}


type Rectangle struct {
	width , height float64
}

type Circle struct {
	x , y , radius float64 
}

func (circle Circle) area() float64{

	return math.Pi * circle.radius * circle.radius
}

func (rect Rectangle) area() float64 {
	return 2 * rect.width * rect.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}






func  main()  {

	circle1:=Circle{radius : 4 , x : 1 , y : 2}
	rectangle1:=Rectangle{width : 34.6 , height : 56.6}
	fmt.Printf("CIRCLE Area: %.3f\n" , getArea(circle1))
	fmt.Printf("RECTANGLE Area: %.3f" , getArea(rectangle1))


	
}