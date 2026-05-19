package main
import "fmt"

func main()  {
	var a int = 10

	var(
		b int = 20
		c string = "Hello"
	) // groyuped variable declaration
	//a:=10 same as var a
   var price , discount float64 = 100.0 , 10.0 // jokhon ek sathe onek variable declare korte hoy tokhon var er por parenthesis er moddhe variable gulo declare kora hoy
   fmt.Println(price)
   fmt.Println(discount)

	if(a==10){
		fmt.Println("Right")
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// explore fmt package , search in golang fmt package in google and see the functions available in fmt package and try to use them in your code