package main

func main() {

}

// Inside the if statement we are creating a new variable called err.
// We are not using the err variable declared as the function return argument.
// The compiler recognizes this and produces the error.
// If the compiler did not report this error,
// you would never see the error that occured inside the if statement.
// The return err variable is what is passed by default

// func returnID() (id int, err error) {
// 	id = 10

// 	if id == 10 {
// 		err := fmt.Errorf("Invalid Id\n")
// 		return
// 	}

// 	return
// }
