package main
import "fmt"

func main(){
var number [6] int = [6] int {12,45,56,77,787,888}
// number := [6] int {12,45,56,77,787,888}
 d := [...]int{1, 2, 3, 4, 5}
fmt.Println(number)// int type ar default value 0, string ar faka string ""
fmt.Println(len(number))
fmt.Println(d)
 for i :=0 ; i<len(number); i++{
	fmt.Println(number[i])
 }
 
 employees := [...]string{"Mobasshir", "Rafiq", "Karim"}
 fmt.Printf("%v\n",employees)

 salary := [...] int { 1,23,4,5,6,6,}
 fmt.Println(salary)

}