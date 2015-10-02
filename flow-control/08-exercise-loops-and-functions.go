package main
import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	i := 0
	n := float64(0)

	for {
		n = z - (z * z - x) / (2 * x);
		if math.Abs(n - z) < 1e-15{
			break
		}
		z = n
		i++;
	}

	return z, i;
}

func main() {
	fmt.Println(Sqrt(4))
}