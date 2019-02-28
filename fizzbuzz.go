package main

import "fmt" // This is the format package, gives input and output.

func main() {
	for i := 1; i <= 100; i++ { // This starts the variable i at 1, does a loop starting at that until it's equal to 100 and then it adds 1 to i each pass.
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
