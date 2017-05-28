package main

import "fmt"

func main() {
	// s := greet("Jay", "Gandhi")
	fmt.Println(greet("Jay", "Gandhi"))

	// Named return - Avoid Using Named return/Naked Return
	fmt.Println(greet1("Jay", "Gandhi"))

	// Multiple Returns
	a, b := greet2("Jay", "Gandhi")
	fmt.Println(a, b)
}

func greet(fname, lname string) string {
	// Printing it to String rather than System Console.
	// Returns string. For diff type it will add space.
	return fmt.Sprint(fname, " ", lname)
}

func greet1(fname, lname string) (s string) {
	// This is what is called a naked return.
	// I have removed the arguments from the return statement.
	// The Go compiler automatically returns the current values in the return arguments local variables.
	s = fmt.Sprint(fname, " ", lname)
	return
}

/*
IMPORTANT
Avoid using named returns.


Occasionally named returns are useful. Read this article for more information:
https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
*/

// Multiple Return
func greet2(fname, lname string) (string, string) {
	return fmt.Sprint(fname, " ", lname), fmt.Sprint(fname, "---", lname)
}
