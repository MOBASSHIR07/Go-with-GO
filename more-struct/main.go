package main

import "fmt"
//  type user struct {
// 	name string
// 	age int
// 	isLoggedIn bool
// 	greet func()
//  }

 type user struct {
	name string
	age int
	isLoggedIn bool
	
 }


func main() {
    // user1 := user {
	// 	name: "jhon",
	// 	age: 25,
	// 	isLoggedIn: false,
		
	// }
	// user1.greet = func() {
	// 	fmt.Println("Hello", user1.name)
	// }

	// user1.greet()

	  user1 := user {
		name: "jhon",
		age: 25,
		isLoggedIn: false,
		
	}
	user1.greet()
	
}
func(u user) greet(){ // go auto user1 k "u" a pass korbe
 fmt.Println("Hello", u.name)
} // receiver  , aivabe kun struct ar method toiri kora jay, it like greet name user struct a akta key ase oi function tai aita ,
//  amra every struct ar (user1 ,2,3..) ar sathe user1.greet() peye jabo
