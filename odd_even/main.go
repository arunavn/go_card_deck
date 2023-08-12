package main

import "fmt"

func main() {
	nSlice := []int{}
	// {1,2,3,4,5,6,7,8,9,10}
	n := 10
	for i := 0; i < n+1; i++ {
		nSlice = append(nSlice, i)
	}
	for _, x := range nSlice {
		if x%2 == 0 {
			fmt.Printf("%v is even \n", x)
		} else {
			fmt.Printf("%v is odd \n", x)
		}
	}
}
