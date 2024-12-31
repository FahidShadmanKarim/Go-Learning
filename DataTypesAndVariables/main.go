package main

import(
	"fmt"
	
)

func main() {

	fmt.Println("Hello World")

	//Types of Integers 
	// int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint
     
    //Ways to decclare variables
    var x int
	x = 10
	fmt.Println(x)

	var y int = 20
	fmt.Println(y)

	//Short hand declaration
	a := 30 // type is inferred
	fmt.Println(a)

	var f32 float32 = 3.14
	var f64 float64 = 3.14 * 2

	fmt.Println(f32)
	fmt.Println(f64)

	String := "Hello World"
	fmt.Println(String)
     
	var b byte = 255
	fmt.Println(b)

	var r rune = 'A'
	fmt.Println(r)

	var bool flag = true
	fmt.Println(bool)


}

