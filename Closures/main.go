package main

import(
	"fmt"
)

func createCounter() func() int {
	count := 0
	increment := func() int {
		count ++
		return count
	}
	return increment
}

func genFib() func() int {
	f1,f2 := 0 , 1
	return func() int {
		f1,f2 = f2,f1 + f2
		return f1
	}
}

func adder() func(int) int {
	 
	sum := 0 
	return func(x int) int{
		sum += x
		return sum
	}
}


func main(){

	
     
	fmt.Println("Closures in Go")
	counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())

	
    
	fib := genFib()

	for i := 0 ; i < 5 ; i++ {
	      fmt.Println(fib())	
	}

	pos,neg := adder(),adder()
	for i := 0 ; i < 10 ; i++ {	
		fmt.Println(pos(i),neg(-2*i))
	}

}