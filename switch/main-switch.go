package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " as : ")
	switch i {
	case 1:
		fmt.Println("one ")
	case 2:
		fmt.Println(" two")
	}
	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday:
		fmt.Println("this is weekend")
	default:
		fmt.Println("this is weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("is am")
	default:
		fmt.Printf("is pm")

	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("im bool")
		case int:
			fmt.Println("im int")
		default:
			fmt.Printf("dont know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(10)
	whatAmI("ffhg")

}
