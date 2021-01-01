// On the top of the file we define the project name
// that the file is associated to

// Here we have one file, but if our app had multiple files, others
// would start with package main too

// If we want to run the app it should be called main and have a function main within it
// If we want to write some reusable package the we would name it differently
// As in our case we want the code to run we just name it package main and include func main

package main

// Package from standard go library
// that is really useful for printing purposes
// there might be more though, so far I have used it just for printing
import "fmt"

// Main function that will be initialised when we run go run main.go
func main() {
	// Declatration of slice that will have elements of type int
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// We are using for loop to loop over the slice that we declared above

	// Syntax might look tricky at first, but here are core things :
	// _ (stands for index, if you need it you can swap _ for "i" or "index" and log it out with fmt.Println() function)
	// _ is used in cases where we know that second variable is available, but is not being used
	// range means "we will go over this slice
	// rest should be reasonably familiar
	for _, n := range s {
		// Her we are using modulus which check what the reminder is if we dived n by 2
		// If it is 0, then clearly it is an even number, else it is odd
		if n%2 == 0 {
			fmt.Println(n, " is even")
		} else {
			fmt.Println(n, " is odd")
		}
	}
}
