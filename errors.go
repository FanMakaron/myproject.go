package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

//func (e ErrNegativeSqrt) String() string {
//	return fmt.Sprint(float64(e))
//}

func (e ErrNegativeSqrt) Error() string {
	//return "cannot Sqrt negative number: " + e
	return "cannot Sqrt negative number: " + fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		var err = ErrNegativeSqrt(x)
		//fmt.Printf(err)
		return x, err
	}

	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
}
