package main
import "fmt"

func main(){
	// arguments are evaluated immediately, function call
	// is not executed until the surrounding function returns
	defer fmt.Println("world")

	fmt.Println("Hello")
}