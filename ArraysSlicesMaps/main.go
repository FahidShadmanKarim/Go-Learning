package main

import(
	"fmt"
)

func main() {
   
	var a [5]int
	
	a[0] = 1
	a[1] = 2
	a[3] = 4

	fmt.Println(a) // [ 1 2 0 4 0 ]
    
	//Iterating over the array using for loop
	for i := 0 ; i < len(a) ; i++{
		fmt.Println(a[i])
	}
    
	//key value pair
	for i, v := range a{
		fmt.Printf("Index %d, Value %d\n", i, v)
	}
    
	// if I dont want the index,but only the value
	for _, v := range a{
		fmt.Printf("Value %d\n", v)
	}

	b := [5]int{1,2,3,4,5}
	fmt.Println(b)

	stringArray := [4]string{
		"Learning",
		"Go",
		"Programming",
		"Language",
	}

	fmt.Println(stringArray)

	// Declaring an array wiothout specifying the size
	c := [...]int{1,2,3,4,5}

	fmt.Println(c)
	
	d := c // d is a copy of c, but doesnt point to the same memory location

	d[0] = 100

	fmt.Println(c)
	fmt.Println(d)

	var matrix [3][3]int
	matrix[0] = [3]int{1,2,3}
	matrix[1] = [3]int{4,5,6}
	matrix[2] = [3]int{7,8,9}

	for i := 0 ; i < 3 ; i++{
		for j := 0 ; j < 3 ; j++{
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	//Slices
	//Slices are like dynamic arrays with no fixed size
   
	arrayA := [5]int{1,2,3,4,5}
	sliceA := arrayA[1:3] // [2 3]
	sliceB := arrayA[2:] // [ 3 4 5]
	sliceC := arrayA[:3] // [ 1 2 3]
	fmt.Println(sliceA, sliceB, sliceC)	

	//Appending to a slice	
	sliceD := []int{1,2,3,4,5}
	sliceD = append(sliceD, 6)
	fmt.Println(sliceD)

	//Appending one slice to another	
	sliceE := []int{7,8,9}
	sliceD = append(sliceD, sliceE...)
	fmt.Println(sliceD)

	//Deleting from a slice 
	//So if you want to delete ith index take the slice form 0 to i-1 and append the slice from i+1 to end
	sliceD = append(sliceD[:2], sliceD[3:]...)
	fmt.Println(sliceD)

	//Copying a slice
	sliceF := make([]int, len(sliceD))
	copy(sliceF, sliceD)
	fmt.Println(sliceF)

	tempSlice := make([]int, 5)// [0 0 0 0 0]
	fmt.Println(tempSlice)

	//Creating a slice with make
	//The capacity of a slice is the maximum number of elements it can hold before it needs to grow.”
	//If you keep adding elements to a slice and it surpasses its current capacity
	//Go will automatically create a larger array, copy the elements over, and use that new array as the slice's underlying array
	sliceG := make([]int, 5, 10) 
	fmt.Println(sliceG)

	//Maps
	//Maps are key value pairs
	//Maps are like dictionaries in python
	//Maps are unordered
	//Maps are reference types
	

	//Creating a map
    mp := make(map[string]int)
	mp["a"] = 1
	mp["b"] = 2
	fmt.Println(mp["a"])

	//Deleting a key value pair
	delete(mp, "a")
	fmt.Println(mp)

	//Checking if a key exists
	_, ok := mp["a"]
	if ok{
		fmt.Println("Key exists")
	}else{
		fmt.Println("Key does not exist")
	}

	mp["a"] = 1

	for key, value := range mp{
		fmt.Printf("Key %s, Value %d\n", key, value)
	}
    
	//Checking if a key exists
	if value, ok := mp["a"]; ok{
		fmt.Println(value)
	}
    
	//Checking if a key exists
	if _, ok := mp["a"]; ok{
		fmt.Println("Key exists")
	}else{
		fmt.Println("Key does not exist")
	}


















}