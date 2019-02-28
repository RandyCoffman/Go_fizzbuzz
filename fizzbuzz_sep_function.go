package main

import "fmt" // This is the format package, gives input and output.
// same code as fizzbuzz1.go, but I'm making the main function cleaner
func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return "FizzBuzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprint(i)
	}

}
