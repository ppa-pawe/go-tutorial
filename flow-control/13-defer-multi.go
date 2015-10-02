package main
import "fmt"

func main() {
	fmt.Println("counting")

	// function calls get stacked and LIFO processed
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}