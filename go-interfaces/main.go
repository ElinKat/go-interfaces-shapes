package main

import "fmt"

func main() {

	shapes := []shape{
		square{length: 15.2},
		square{length: 4.9},
		circle{radius: 7.5},
		circle{radius: 12.3},
	}

	for _, value := range shapes {
		printShapeInfo(value)
		fmt.Println("---")
	}
}
