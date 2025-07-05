// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func pointerBehavior(x *int) (int, error) {
	fmt.Println("Inside Function")
	fmt.Println("----------------------------------------")

	// Check if calling parameter is not null.
	if !(x == nil) {
		*x = int(math.Pow(float64(*x), 2))
		fmt.Println("Param Mem Addr: [", x, "]\nLocal Mem Addr: [", &x, "]\nLocal Value   : [", *x, "]")
		fmt.Println("----------------------------------------")
		return *x, nil
	} else {
		fmt.Println("----------------------------------------")
		return 0, errors.New("no valid memory address")
	}
}

func main() {
	// Variable declarations.
	var a int = 12
	var b *int

	// Not nil variable.
	fmt.Println("Not Nil Test Call")
	fmt.Println("----------------------------------------")
	fmt.Println("Before Call")
	fmt.Println("Variable Mem Addr [a] is [", &b, "]\nVariable Value    [a] is [", a, "]")
	fmt.Println()
	fmt.Println("========================================")
	a, err := pointerBehavior(&a)
	fmt.Println("========================================")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println()
	fmt.Println("After Call")
	fmt.Println("----------------------------------------")
	fmt.Println("Variable Mem Addr [a] is [", &b, "]\nVariable Value    [a] is [", a, "]")
	fmt.Println("----------------------------------------")

	// Nil variable.
	fmt.Println()
	fmt.Println("Nil Test Call")
	fmt.Println("----------------------------------------")
	fmt.Println("Before Call")
	fmt.Println("Variable Mem Addr [b] is [", b, "]")
	fmt.Println()
	fmt.Println()
	fmt.Println("========================================")
	ab, err := pointerBehavior(b)
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("After Call")
	fmt.Println("----------------------------------------")
	if err != nil {
		fmt.Println("Variable Mem Addr [b] is: [", ab, "]")
		fmt.Println("Error Message.          : [", err, "]")
		os.Exit(1)
	}
	fmt.Println("----------------------------------------")
	fmt.Println()

}