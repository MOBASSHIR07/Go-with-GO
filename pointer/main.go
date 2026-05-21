package main 
import "fmt"
 func change (x *int){
   
	 *x = 1000000
	 fmt.Println("inside, value x  :",*x)

 }

func main () {
// 	a:=20

// 	p:= &a
// 	a = 43
// 	fmt.Println("Hello")

// 	fmt.Println(&a) // & show memory address
// 	fmt.Println("a:" , a)
// 	fmt.Println("p:" , p)
// 	fmt.Println("p:" , *p) // *p de reference

// 	*p = 10000000 // pointer variable hoye gele aivabe reassign kore
//    fmt.Println("_________________________")
// 	fmt.Println("a:" , a)
// 	fmt.Println("p:" , p)
// 	fmt.Println("p:" , *p) 

//____________________________________________________


  y:= 20
  change(&y)
  fmt.Println(y)
 
}