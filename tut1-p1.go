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
	fmt.Println("======================Exercise 2==========================================")
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
	fmt.Println("======================Exercise 3==========================================")
	//Ex3
	var f, g uint16
	fmt.Printf("Input number c: ")
	fmt.Scanf("%d", &f)
	g = f * 2
	fmt.Printf("Double of input number is: %d\n", g)
	if g > 100 {
		fmt.Println("The result is larger than 100")
	} else {
		fmt.Println("The result is smaller than 100")
	}
	fmt.Println("======================Exercise 4==========================================")
	//Ex4
	var h uint16
	fmt.Printf("Type the number of the month: ")
	fmt.Scanf("%d", &h)
	if h <= 12 && h >= 1 {
		fmt.Printf("Number of the month: %d\n", h)

		if h >= 1 && h <= 3 {
			fmt.Println("Season 1")
		} else if h >= 4 && h <= 6 {
			fmt.Println("Season 2")
		} else if h >= 7 && h <= 9 {
			fmt.Println("Season 3")
		} else if h >= 10 && h <= 12 {
			fmt.Println("Season 4")
		}

	} else {
		fmt.Printf("The input number is not number of the month\n")
	}
	fmt.Println("======================Exercise 5==========================================")
	//Ex5
	var distance, price uint64
	fmt.Printf("Type the distance: ")
	fmt.Scanf("%d", &distance)
	fmt.Printf("The distance is: %dkm\n", distance)

	if distance <= 3 {
		price = distance * 20500
	} else if distance > 3 && distance <= 15 {
		price = distance * 15350
	} else if distance > 15 && distance <= 39 {
		price = distance * 12999
	} else if distance > 39 {
		price = distance * 8000
	}

	if distance%10 == 0 {
		price = price - price*20/100
	}

	fmt.Printf("The price is: %d VND\n", price)

	fmt.Println("======================Exercise 6==========================================")
	//Ex6
	fmt.Println("==========================================================================")
}
