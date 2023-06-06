package main

import (
	"fmt"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}

	if mToSend.recipient.name == "" {
		return false
	}

	if mToSend.sender.number == 0 {
		return false
	}

	if mToSend.recipient.number == 0 {
		return false
	}

	return true
}

// don't touch below this line

func test(mToSend messageToSend) {
	fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
		mToSend.message,
		mToSend.sender.name,
		mToSend.sender.number,
		mToSend.recipient.name,
		mToSend.recipient.number,
	)
	fmt.Println()
	if canSendMessage(mToSend) {
		fmt.Println("...sent!")
	} else {
		fmt.Println("...can't send message")
	}
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		message: "you have an appointment tommorow",
		sender: user{
			name:   "Brenda Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have an event tommorow",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "Suzie Sall",
			number: 0,
		},
	})
	test(messageToSend{
		message: "you have an party tommorow",
		sender: user{
			name:   "Njorn Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Eli Halafax",
			number: 0,
		},
		recipient: user{
			name:   "Whitaker Sue",
			number: 19035558973,
		},
	})
}

// NESTED STRUCTS IN GO
// Structs can be nested to represent more complex entities:
//
// type car struct {
//   Make string
//   Model string
//   Height int
//   Width int
//   FrontWheel Wheel
//   BackWheel Wheel
// }
//
// type Wheel struct {
//   Radius int
//   Material string
// }
// The fields of a struct can be accessed using the dot . operator.
//
// myCar := car{}
// myCar.FrontWheel.Radius = 5
// ASSIGNMENT
// Textio has a bug, we've been sending texts with information missing! Before we send text messages in Textio, we should check to make sure the required fields have non-zero values.
//
// Notice that the user struct is a nested struct within the messageToSend struct. Both sender and recipient are user struct types.
//
// Complete the canSendMessage function. It should return true only if the sender and recipient fields each contain a name and a number. If any of the default zero values are present, return false instead.
