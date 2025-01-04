package main

import(
	"fmt"
)

//When to use pointers
//1. When we need to update 
//2. Pointer = 8bytes
// Pointer Use pointers to reduce memory usage: Pointers can be a useful tool for reducing memory usage in Go. 
//By using pointers to store large data structures, you can avoid copying the entire data structure every time it 
//is passed to a function or stored in a data structure.

type User struct {
	email string
	username string
	age int
}

//8 bytes if pointer passed, else n amount of bytes
//where x = size of User
func (u *User) updateEmail(email string){
	u.email = email
}

func modString(s *string){
	fmt.Println("Inside modString:",s)
	*s = "modified"
}

func main(){
     
	fmt.Println("Pointers")

	
	var x int = 10
	var ptr *int = &x
    var ptr1 *int = ptr
	fmt.Println(*ptr)
	fmt.Println(*ptr1)
	*ptr1 = 12
	fmt.Println(x)
	*ptr = 15
	fmt.Println(x)

	s := "hello"
	fmt.Println("Before:",s)
	modString(&s)
	fmt.Println("After:",s)

	user := User{
		email : "fahid@foo.com",
	}

	fmt.Println(user.email)

	user.updateEmail("fahid@yahoo.com")

	fmt.Println(user.email)
     
    var str string = "Fahid"
	var sPtr *string = &str
	fmt.Println(*sPtr)
	



	











}