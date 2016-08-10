package main

import "fmt"

import "math"

func main() {
	fmt.Println("======================Exercise 1==========================================")
	//Ex1
	var myname string
	fmt.Printf("What is your name? ")
	fmt.Scanf("%s", &myname)
	fmt.Printf("Hello %s\n", myname)
	fmt.Println("======================Excercise 2==========================================")
	//Ex2
	var a, b, c, d, e float64
	fmt.Printf("Input number a & b: ")
	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	c = math.Pow(a, 2)
	d = math.Pow(b, 2)
	e = c + d
	fmt.Printf("Result of square a: %.1f\n", c)
	fmt.Printf("Result of square b: %.1f\n", d)
	fmt.Printf("Result of the math %.1f\n", e)
	fmt.Println("======================Excercise 3==========================================")
	//Ex3
	var f, g uint8
	fmt.Printf("Input number c: ")
	fmt.Scanf("%d", &f)
	g = f * 2
	fmt.Printf("Double of input number is: %d\n", g)
	if g > 100 {
		fmt.Println("The result is larger than 100")
	} else {
		fmt.Println("The result is smaller than 100")
	}

	fmt.Println("================================================================")
}
