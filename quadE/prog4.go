package main

import "fmt"

func main() {

	QuadA(1, 5)
	
}

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	// Print top border
	fmt.Print("A")
	for i := 1; i < x-1; i++ {
		fmt.Print("B")
	}
	if x > 1 {
		fmt.Print("C")
	}
	fmt.Println()
	// Print sides
	for j := 1; j < y-1; j++ {
		fmt.Print("B")
		for i := 1; i < x-1; i++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Print("B")
		}
		fmt.Println()
	}
	// Print bottom border
	if y > 1 {
		fmt.Print("C")
		for i := 1; i < x-1; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Print("A")
		}
		fmt.Println()
	}

}
