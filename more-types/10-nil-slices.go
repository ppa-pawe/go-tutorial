package main
import "fmt"


func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))

	z = append(z, 1)

	fmt.Println(z, len(z), cap(z))

	if z == nil {
		fmt.Println("nil!")
	}
}