package main

import "fmt"

func main() {
	var max int

	fmt.Printf("Enter the number of elements: ")
	fmt.Scanln(&max)
	number := make([]int, max)
	fmt.Printf("Enter the numbers with space separated to find the multiplication table of it: ")
	for i := 0; i < max; i++ {
		fmt.Scan(&number[i])
	}
	fmt.Printf("Enter the number range up to which to calculate multiples of %d : ", number)
	fmt.Scan(&max)
	for j := number[0]; j <= len(number)-1; j++ {
		for i := 1; i <= max; i++ {
			fmt.Printf("%d X %d = %d \n", number[j], i, i*number[j])
		}
	}
}

// package main

// import "fmt"

// func main() {
// 	X := make([]int, 3)
// 	fmt.Scanln(&X[0], &X[1], &X[2])
// 	fmt.Printf("%v\n", X)
// }
