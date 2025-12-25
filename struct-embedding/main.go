// package main

// import "fmt"

// type base struct {
// 	num int
// }

// func (b base) describe() string {
// 	return fmt.Sprintf("base with number=%v", b.num)
// }

// type container struct {
// 	base
// 	str string
// }

// func main() {
// 	c := container{
// 		base: base{
// 			num: 10,
// 		},
// 		str: "this is some string",
// 	}
	
// 	fmt.Printf("c={num: %v, str: %v}\n", c.base.num, c.str)
// 	fmt.Println(c.base.num)
// 	fmt.Println(c.num)
	
// 	fmt.Println("describe:", c.describe())
	
	
// 	type describer interface {
// 		describe() string
// 	}
	
// 	var d describer = c
// 	fmt.Println("describer is:", d.describe(), )
// }