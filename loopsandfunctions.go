package main
import (
	"fmt"
	"math"
)
func Sqrt(x float64) (float64) {
	z := 1.0
	delta := 1e-12
	counter := 0
	
	for{
		temp:= z
		z = z - ((math.Pow(z,2)) - x)/(2 * z)
		if math.Abs(z-temp) <= delta{
			break
		}
		counter++
	}
	fmt.Println("Counter:", counter)
	return z
}
func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
}

