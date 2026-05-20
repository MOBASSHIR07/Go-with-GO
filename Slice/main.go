package main
import "fmt"
func main(){
	var order = [6] int {1,3,5,76,8,8}
	slice:=order[0:4] // 0 index to 4-1 index
	//slice:=order[2:] // 2 index to last
	//slice:=order[:] //  all element 
	// slice ager array k change kore dey
	// slice len , capacity , pointer thake

	fmt.Println(slice)
}