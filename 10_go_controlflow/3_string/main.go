package main

import "fmt"

func main() {
	i := 10
	fmt.Println(i, " Difference between ", 'i')
	// first i is integer value you've set
	// 'i' 105 is ascii/unicode value of letter i
	foo := 'a'
	fmt.Println(foo)              // 97
	fmt.Printf("%T\n", foo)       // int32
	fmt.Printf("%T\n", rune(foo)) // int32
	// if you've declared foo := "a" then you can't convert into rune
	// Row String literal -> `Jay` (backticks)
	// interpreted string literal -> "Jay" Double Quotes
	// You can easily print with Raw string literal
	fmt.Println(`
			<html>
				<head>
					<title>Jay Gandhi"</title>
				</head>
				<body>
					etc
				</body>
			</html>
	`)
}
