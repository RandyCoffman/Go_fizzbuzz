package main

import "fmt" // This is the format package, gives input and output.
// Used switches to make the conditions better, trying to get a feel of the language. Going to try to do a bigger project soon.
func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(i int) string {
	switch {
	case i%3 == 0 && i%5 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprint(i)

	}

}
