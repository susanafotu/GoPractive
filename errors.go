package main


import (
	"fmt"
)


type ErrNegativeSqrt float64


func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	return x, nil
}


func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v", (float64(e)))
}


func main() {
	x, err := Sqrt(2)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
	
	x, err = Sqrt(-2)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}

