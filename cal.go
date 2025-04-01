package main

import (
	"fmt"
	"math"
)

func sum() {
	var a, b int
	fmt.Print("enter the first number:")
	fmt.Scan(&a)
	fmt.Print("enter the second number:")
	fmt.Scan(&b)
	c := a + b
	fmt.Printf("The sum of %v and %v is %v", a, b, c)
}

func sub() {
	var a, b int
	fmt.Print("enter the first number: ")
	fmt.Scan(&a)
	fmt.Print("enter the second number: ")
	fmt.Scan(&b)
	c := a - b
	fmt.Printf("The Subtraction of %v and %v is %v", a, b, c)
}

func prod() {
	var a, b int
	fmt.Print("enter the first number: ")
	fmt.Scan(&a)
	fmt.Print("enter the second number: ")
	fmt.Scan(&b)
	c := a * b
	fmt.Printf("The Product of %v and %v is %v", a, b, c)
}

func div() {
	var a, b int
	fmt.Print("enter the first number: ")
	fmt.Scan(&a)
	fmt.Print("enter the second number: ")
	fmt.Scan(&b)
	c := a / b
	fmt.Printf("The Division of %v and %v is %v", a, b, c)
}

func sqr() {
	var a float64
	fmt.Print("enter the number: ")
	fmt.Scan(&a)
	c := math.Pow(a, 2)
	fmt.Printf("The square of %v is %v", a, c)
}

func sqrr() {
	var a float64
	fmt.Print("enter the number: ")
	fmt.Scan(&a)
	c := math.Sqrt(a)
	fmt.Printf("The square root of %v is %v", a, c)
}

func sign() {
	var a float64
	fmt.Print("enter the angle in degrees: ")
	fmt.Scan(&a)
	b := a * (math.Pi / 180)
	c := math.Sin(b)
	d := math.Asin(c) * (180 / math.Pi)
	fmt.Printf("The sine of angle %.1f is %.1f \n", a, c)
	fmt.Printf("The Arcsin of %.1f is %.1f", c, d)
}
func cus() {
	var a float64
	fmt.Print("enter the angle in degrees: ")
	fmt.Scan(&a)
	b := a * (math.Pi / 180)
	c := math.Cos(b)
	d := math.Acos(c) * (180 / math.Pi)
	fmt.Printf("The Cosine of angle %.1f is %.1f \n", a, c)
	fmt.Printf("The Arccosine of %.1f is %.1f", c, d)
}
func tann() {
	var a float64
	fmt.Print("enter the angle in degrees: ")
	fmt.Scan(&a)
	b := a * (math.Pi / 180)
	c := math.Tan(b)
	d := math.Atan(c) * (180 / math.Pi)
	fmt.Printf("The Tangent of angle %.1f is %.1f \n", a, c)
	fmt.Printf("The Arctan of %.1f is %.1f", c, d)
}

func main() {
	var option int

	fmt.Println("\n\n===== Simple calculator ====== \n")
	fmt.Println("=== Options ===")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Product")
	fmt.Println("4. Division")
	fmt.Println("5. Square")
	fmt.Println("6. Square root")
	fmt.Println("7. Sine and Arcsin")
	fmt.Println("8. Cosine and Arccosine")
	fmt.Println("9. Tagent and Arctan")
	fmt.Println("10. Exit Calculator \n")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&option)
	if option == 1 {
		sum()
	} else if option == 2 {
		sub()
	} else if option == 3 {
		prod()
	} else if option == 4 {
		div()
	} else if option == 5 {
		sqr()
	} else if option == 6 {
		sqrr()
	} else if option == 7 {
		sign()
	} else if option == 8 {
		cus()
	} else if option == 9 {
		tann()
	} else if option == 10 {
		fmt.Println("Exiting the calculator")
		return
	}

}
