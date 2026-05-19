package main
import "fmt"
 func main()  {

	fmt.Println("welcome to the application")
	
	// get user name as input
	var name string 
	fmt.Println("Enter your name :") 
	fmt.Scanln(&name) // & is the pointer of the variable name
	fmt.Println("My name is :" ,name)
	var age int
		fmt.Println("Enter your age :")
	fmt.Scanln(&age)
	fmt.Println("age is", age)
 }