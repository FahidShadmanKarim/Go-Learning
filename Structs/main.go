package main

import "fmt"

type Address struct{
	city string
	state string
}

type Person struct {

	name string
	age  int
	address Address

}

type Employee struct {
    
	firstName string
	lastName  string
	age 	  int
	

}

type image struct{

	data map[int]int
}


func main(){
   
	emp1 := Employee{
		firstName: "Fahid",
		lastName:  "Shadman",
		age:		25,
	}

	emp2 := Employee{"Thomas", "Paul", 29}
    
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// you can create anonymous structs as well
	emp3 := struct{
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
	}

	fmt.Println(emp3)

	// The dot operator is used to access fields
	fmt.Println(emp1.firstName)


	p := Person{
		name :"Alex",
		age: 30,
		address: Address{
			city : "Chicago",
			state : "Illinois",
		},

	}
	p1 := Person{
		name :"Alex",
		age: 30,
		address: Address{
			city : "Chicago",
			state : "Illinois",
		},

	}

	fmt.Println("Address: ",p.address)

	if p == p1 {
		fmt.Println("Two persons are identical")
	}else{
		fmt.Println("They are not identical")
	}

	image1 := image{
		data: map[int]int{
			0:255,
	}}

	image2 := image{
		data: map[int]int{
			0:255,
	}}

	// we cant compare image1 and image2 becase
	// their fields which is a map cant be compared
    
	fmt.Println(image1)
	fmt.Println(image2)


	












}