package basics1

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) area() string {
	return fmt.Sprintf("The area of the Rectangle is %d", r.height*r.width)
}

func structMethods() {
	rectangle1 := Rectangle{
		width:  5,
		height: 6,
	}

	area1 := rectangle1.area()

	fmt.Println(area1)
}
