package main

import "fmt"

func main() {

	var p *int      // "p" is a pointer to a variable that will be an integer type

	i := 42         // shorthand definition of variable "i"

	p = &i          // The "&" operator generates a pointer to its operand.

	fmt.Println(p)  // returns memory address, like (0x20e010)

			// The `*` operator denotes the point's underlying value
	fmt.Println(*p) // read i through the pointer p (de-reference operation)

	*p = 21         // set i through the pointer p  (de-reference operation)

	fmt.Println(i)       // i is now 21
	fmt.Println(*p)      // p points to i, it can't help but be 21
	fmt.Println(p)       // this shows the VALUE of i's memory
        fmt.Println(&i)      // show the value of i's memory
	
	fmt.Println(&p)

	var z *int

	z = &i

	*z = 418
	
	fmt.Println(i)      // i is now 418
	fmt.Println(*p)     // *p should also show i is 418
	fmt.Println(*z)     // *z should also show i is 418
}
