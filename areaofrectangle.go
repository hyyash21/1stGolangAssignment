package main

import "fmt"

func main() {
	var length, breadth, area, peri float64
	fmt.Println("Enter the Length in cm: ")
	fmt.Scanln(&length)
	fmt.Println("Enter the Breadth in cm: ")
	fmt.Scanln(&breadth)
	area = length * breadth
	peri = 2*(length+breadth)
	fmt.Printf("Area of Rectangle: %.1f cm\u00B2\n", area)
	fmt.Printf("Perimeter of Rectangle: %.1f cm\u00B2\n", peri)

}
