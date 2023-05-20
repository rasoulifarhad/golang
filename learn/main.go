package main

import (
	"fmt"
  	"strings"
        "sample.com/learn/math"
	"github.com/pborman/uuid"
        "sample.com/learn/math/advanced"
)

func init(){
	fmt.Println("In main init")
}

func main() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)


       fmt.Println(math.Add(2, 1))
       fmt.Println(math.Subtract(2, 1))
       fmt.Println(advanced.Mul(2,3))
}



func add(a, b int) int {
    return a + b
}
