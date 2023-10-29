package main

import "fmt"

func main() {

	QuadA(1, 1)
	
}

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	// Print top border
	fmt.Print("/")
	for i := 1; i < x-1; i++ {
		fmt.Print("*")
	}
	if x > 1 {
		fmt.Print("\\")
	}
	fmt.Println()
	// Print sides
	for j := 1; j < y-1; j++ {
		fmt.Print("*")
		for i := 1; i < x-1; i++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// Print bottom border
	if y > 1 {
		fmt.Print("\\")
		for i := 1; i < x-1; i++ {
			fmt.Print("*")
		}
		if x > 1 {
			fmt.Print("/")
		}
		fmt.Println()
	}

}
