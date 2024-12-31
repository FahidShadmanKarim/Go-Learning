package main

import(
	"fmt"
)

func main() {
 
	x := 5
	y := 10
    

	//If else
	if x > y {
		fmt.Println("x is greater than y")
	} else{
		fmt.Println("y is greater than x")
	}

	//for loop
	for i := 0; i < 10 ; i++{
		fmt.Println(i)
	}
    
	//while loop like
	j := 0
	for j < 10{
		fmt.Printf("j = %d\n", j)
		j++
	}

	//switch case

	for i := 1 ; i < 5 ; i++{
		switch i {
			case 1:
				fmt.Println("i is 1")
			case 2:
				fmt.Println("i is 2")
			case 3:
				fmt.Println("i is 3")
			default:
				fmt.Println("i is greater than 3")
		}
	}

	






	


}