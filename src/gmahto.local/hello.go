package main

import "fmt"

type salutation struct {
	name     string
	greeting string
}

func test(message *salutation, a int, b int) {

	a = 11
	b = 13
	// message.name = "Gaurav"
	// message.greeting = "Hello!"

	*message = salutation{name: "Gaurav", greeting: "Hello!"}

	// fmt.Println(*message, a, b)

}

func main() {

	var message = salutation{"User", "Hello GO"}
	a, b := 1, 2 // Declaration and assignment.

	fmt.Println("Before calling #test()")

	fmt.Println(message.greeting, message.name, a, b)

	fmt.Println("After calling #test()")

	test(&message, a, b)

	fmt.Println(message.greeting, message.name, a, b)

}
