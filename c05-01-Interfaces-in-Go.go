// INTERFACES IN GO
// Interfaces are collections of method signatures. A type "implements" an interface if it has all of the methods of the given interface defined on it.
//
// In the following example, a "shape" must be able to return its area and perimeter. Both rect and circle fulfill the interface.
//
// type shape interface {
//   area() float64
//   perimeter() float64
// }
//
// type rect struct {
//     width, height float64
// }
// func (r rect) area() float64 {
//     return r.width * r.height
// }
// func (r rect) perimeter() float64 {
//     return 2*r.width + 2*r.height
// }
//
// type circle struct {
//     radius float64
// }
// func (c circle) area() float64 {
//     return math.Pi * c.radius * c.radius
// }
// func (c circle) perimeter() float64 {
//     return 2 * math.Pi * c.radius
// }
// When a type implements an interface, it can then be used as the interface type.
//
// ASSIGNMENT
// The birthdayMessage and sendingReport structs have already implemented the getMessage methods. The getMessage method simply returns a string, and any type that implements the method can be considered a message.
//
// First, add the getMessage() method as a requirement on the message interface.
//
// Second, complete the sendMessage function. It should print a message's message, which it obtains through the interface method. Notice that your code doesn't need to worry at all about whether a specific message is a birthdayMessage or a sendingReport!

package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("====================================")
}

func main() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "Bill Deer",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})
}
