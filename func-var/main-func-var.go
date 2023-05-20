package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums, " ")
	tot := 0
	for _, num := range nums {
		tot += num
	}
	fmt.Println(tot)
}
func main() {
	sum(1, 2)
	sum(3, 7, 7)
	nums := []int{9, 9, 9}
	sum(nums...)

}
