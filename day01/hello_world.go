package main

import (
	"fmt"
	"os"
	"strconv"

	"golang-tutorial-project/day01/myvars"
)

func main() {
	fmt.Println("Hello, World!")

	var a = 10
	fmt.Println(a)

	b := 20

	fmt.Println(a + b)

	var c int
	fmt.Println(c)

	var d float32
	fmt.Println(d)

	var e string
	fmt.Println(e)

	var f bool
	fmt.Println(f)

	var x int = 10
	var y int = 20

	fmt.Println(x + y)

	var z int = x + y
	z = z + 10
	fmt.Println(z)

	const PI float32 = 3.14
	fmt.Println(PI)

	// PI = 3.1678
	// fmt.Println(PI)

	// loop starts here
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	{
		xx := 10
		fmt.Println(xx)
	}

	// fmt.Println(i)
	// fmt.Println(xx)

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println(i)

	fmt.Println("==========Break==========")

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("==========Continue==========")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("==========Continue with even numbers==========")
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}

	ab := 10

	if ab%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	a, b = 10, 20
	if a > b {
		fmt.Println("a is greater")
	} else if a < b {
		fmt.Println("b is greater")
	} else {
		fmt.Println("Both are equal")
	}

	// switch case
	fmt.Println("==========Switch Case==========")
	dayStr := os.Args[1]

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		fmt.Println("Invalid day")
		return
	}

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	// transform to if else
	if day == 1 {
		fmt.Println("Monday")
	} else if day == 2 {
		fmt.Println("Tuesday")
	} else if day == 3 {
		fmt.Println("Wednesday")
	} else {
		fmt.Println("Invalid day")
	}

	fmt.Println(myvars.PI)
}
