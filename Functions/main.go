package main

import(
	"fmt"
)

func add( x int , y int ) int {
	return x + y
}

func sumOfAList(array []int)int{
	sum := 0 
	for i := 0 ; i < len(array) ; i++{
		sum += array[i]
	}
	return sum
}

func mutiValueReturn() (int , int){
	return 5,6
}

func factorial(n int) int {
	if n == 0 {
		return 1 
	}
	return n * factorial( n - 1 )
}

func mx( a int , b int , args...int ) int {
      
	mx := a
	if b > mx {
		mx = b
	}
	for i := 0 ; i < len(args) ; i++{
		if args[i]  > mx {
			mx = args[i]
		}
	}
	return mx
}

func min( a int , b int , args...int)int{
	min := a
	if b < min {
		min = b
	}
	for i := 0 ; i < len(args) ; i++{
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}

func function1( anonymous func(p,q string) string ){
    //takes an anonymous function as a parameter
	fmt.Println(anonymous("My name","is"))
}

func main() {
     
	fmt.Println(add(2,3))

	fmt.Println(sumOfAList([]int{1,2,3,4,5}))
	a , b := mutiValueReturn()

	fmt.Println(a,b)
	fmt.Println(factorial(5))
	fmt.Println(mx(1,2,3,4,5),min(1,2,3,4,5))
    

	//You can assign func to a variable
	name := func(p,q string) string {
		return p + " " + q + " " + "Fahid"
	}
    
	function1(name)

	//Passing parameters to an anonymous function
	func(s string){
		fmt.Println(s)
	}("Hello World")
 
	swap := func( x , y int) ( int , int) {
		return y , x
	}

	x,y := 5,10
	fmt.Println(x,y)
	x,y = swap( x,y )
	fmt.Println(x,y)

	//To be studied : Closure
	






	
	
}