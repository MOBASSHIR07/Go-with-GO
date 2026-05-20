package main
import "fmt"
func main(){
	var order = [6] int {1,3,5,76,8,8}
	slice:=order[0:4] // 0 index to 4-1 index
	//slice:=order[2:] // 2 index to last
	//slice:=order[:] //  all element 
	// slice ager array k change kore dey
         //slice শুরু হওয়ার index থেকে original array এর শেষ পর্যন্ত কয়টা element available।
	// slice len , capacity , pointer thake
	//কতগুলো element currently আছে।
	//Go তে slice change করলে original array ও change হয় — কারণ slice copy না, underlying array কে reference করে।
	//Slice নিজে dynamic মনে হলেও আসলে এটা fixed-size underlying array এর উপর reference হিসেবে কাজ করে, আর capacity শেষ হলে Go নতুন বড় array বানিয়ে data copy করে বলে slice dynamically grow করতে পারে।
	//Slice এ নতুন element add করতে সাধারণত append() ব্যবহার করা হয়
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	fmt.Println(slice)
}