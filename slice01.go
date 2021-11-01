/* Alta3 Research - Slice without pre-existing array */

package main

import "fmt"

func main() {

        // Slice - without pre-existing array              
    	mySlice3 := make([]int, 4, 4)      // mySlice has lenght and 
                                           // capacity 5, and it's 
                                           // filled with zero values 
                                           // for int, so it's [0 0 0 0]
                                           
    	mySlice4 := make([]int, 4)         // is the same as mySlice3

    	fmt.Println("Contents of mySlice3:", mySlice3)
    	fmt.Println("Contents of mySlice4:", mySlice4)

}

