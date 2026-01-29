The interface{} type in Go is often called the empty interface. It takes any value.

Every interface value in Go is made of two parts. One holds the type, and the other holds a pointer to the data. 
When the interface is declared with methods, it also includes a method table, but for interface{} that’s not needed. 


Type assertions
A type assertion provides access to an interface value's underlying concrete value.

t := i.(T)
This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

t, ok := i.(T)
If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

package main

import "fmt"

func main() {
	var i interface{} = "hi"

	if s,ok := i.(string) ; ok {
		fmt.Printf("%s\n", s)
	} else {
		fmt.Printf("XXX\n")
	}
		
}


Type switches
A type switch is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

package main

import "fmt"

type Profile struct {
	ID int
}

func main() {
	var v interface{} = 5

	// 1. Single Type Assertion
	if i, ok := v.(int); ok {
		fmt.Printf("Single assertion: %d\n", i)
	}

	// 2. The Type Switch
	switch i := v.(type) {
	case string:
		// i is a string here
		fmt.Printf("It's a string: %s\n", i)
	case int:
		// i is an int here
		fmt.Printf("It's an int: %d\n", i)
	case Profile:
		// i is a Profile struct here
		fmt.Printf("It's a Profile with ID: %d\n", i.ID)
	default:
		fmt.Printf("Unknown type: %T\n", i)
	}
}


The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.

This switch statement tests whether the interface value i holds a value of type T or S. In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i. In the default case (where there is no match), the variable v is of the same interface type and value as i.
