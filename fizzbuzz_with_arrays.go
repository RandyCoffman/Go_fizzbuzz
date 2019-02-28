package main

import "fmt"

func main() {
	arr := makeArr()
	for value := range arr {
		fmt.Println(fizzbuzz(value))
	}
}

func makeArr() [101]int {
	var arr [101]int
	for index := range arr {
		arr[index] = index
	}
	return arr
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
