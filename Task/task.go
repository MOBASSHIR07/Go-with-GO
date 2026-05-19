package main
import "fmt"

func coffeeOrder(cName string , coffeeType string , price int)  {
	fmt.Printf("Order from %s : %s coffee cost %d taka \n", cName, coffeeType, price)
}

func main(){
	fmt.Println("Hello, World!")
	coffeeOrder("Alice", "Latte", 150)
	coffeeOrder("Bob", "Espresso", 120)
}