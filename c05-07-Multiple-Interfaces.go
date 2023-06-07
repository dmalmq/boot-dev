package main

import (
	"fmt"
)

func (e email) cost() float64 {
	if e.isSubscribed == true {
		return 0.01 * float64(len(e.body))
	} else {
		return 0.05 * float64(len(e.body))
	}
}

func (e email) print() {
	fmt.Printf(e.body)
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test(e, e)
}

// MULTIPLE INTERFACES
// A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements.
//
// ASSIGNMENT
// Add the required methods so that the email type implements both the expense and printer interfaces.
//
// COST()
// If the email is not "subscribed", then the cost is 0.05 for each character in the body. If it is, then the cost is 0.01 per character.
//
// PRINT()
// The print method should print to standard out the email's body text.
