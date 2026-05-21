package main

import "fmt"

type user struct {
	name  string
	email string
}

func main() {

	john := user{
		name:  "john",
		email: "sadik@gmail.co",
	}

	rahim := user{"rahim", "rahim@gmail.com"} // positional way

	fmt.Println(john)

	fmt.Printf("%+v\n", john)

	fmt.Println(rahim)
}
// struct এটা basically multiple related data একসাথে রাখার way