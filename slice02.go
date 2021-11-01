/* Alta3 Research - Slice and append function */

package main

import "fmt"

func main() {

        // append an element to a slice
        a := []int{1, 2}
        a = append(a, 3) // a == [1 2 3]
        fmt.Println(a)   // [123]

        // Concatenate two slices
        b := []int{1, 2}
        c := []int{11, 22}
        b = append(b, c...) // b == [1 2 11 22]
        fmt.Println(b)
        
        // The result does not depend on whether the arguments overlap
        // so we can concatenate a slice with itself
        z := []int{1, 2}
        z = append(z, z...) // z == [1 2 1 2]
        fmt.Println(z)
        
}

