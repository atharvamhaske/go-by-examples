package main

import "fmt"

type Shape interface {
	Area() int
} //define a interface this will be like a cotract

type Dimensions struct {
	width int
	height int
} //base struct (this will be embedded)

func (d Dimensions) Area() int {
	return d.width * d.height
} //method defined on base struct


type Rectangle struct {
	Dimensions
}
//jo bhi Area() int provide karega wo by default ek shape hi maana jayega, easy term: jiska bhi Area calculate hoke int me ayega wo apne
// aap shape maana jayega
func main() {
	r := Rectangle{
		Dimensions{
			width: 10,
			height: 20,
		},
	}
	
	fmt.Println("Area of rectangle is:", r.Area())
}

//Rectangle implements Shape because it has Area()