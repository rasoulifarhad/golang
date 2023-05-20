package math

import "fmt"
import "sample.com/learn/math/advanced"

func init(){
	fmt.Println("In Mul math  init")
}

func Mul(a, b int) int {
    return advanced.Mul(a , b)
}

