package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		authI.username,
		authI.password,
	)
}

// ?

// don't touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}

// STRUCT METHODS IN GO
// While Go is not object-oriented, it does support methods that can be defined on structs. Methods are just functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function.
//
// type rect struct {
//   width int
//   height int
// }
//
// // area has a receiver of (r rect)
// func (r rect) area() int {
//   return r.width * r.height
// }
//
// r := rect{
//   width: 5,
//   height: 10,
// }
//
// fmt.Println(r.area())
// // prints 50
// A receiver is just a special kind of function parameter. Receivers are important because they will, as you'll learn in the exercises to come, allow us to define interfaces that our structs (and other types) can implement.
//
// ASSIGNMENT
// Let's clean up Textio's authentication logic. We store our user's authentication data inside an authenticationInfo struct. We need a method that can take that data and return a basic authorization string.
//
// The format of the string should be:
//
// Authorization: Basic USERNAME:PASSWORD
// Create a method on the authenticationInfo struct called getBasicAuth that returns the formatted string.
