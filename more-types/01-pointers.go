package main
import "fmt"

func main(){
	i, j := 42, 701

	p := &i

	// prints 42
	fmt.Println(*p)

	*p = 21
	// prints 21
	fmt.Println(i)

	p = &j
	*p = *p / 37

	// prints 701/37
	fmt.Println(j)
}