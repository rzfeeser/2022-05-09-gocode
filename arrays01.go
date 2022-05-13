/* RZFeeser | Alta3 Research
   Array basics*/
   
package main

import "fmt"

func main() {

    // Create an array that holds 5 integers
    var scores [5]int
    fmt.Println("emp:", scores)

    // set a value at a particular index
    scores[4] = 100
    i := scores[4]

    fmt.Println("set:", scores)       // display all scores
    fmt.Println("get:", scores[4])    // display the value @ index 4

    scores[4] = 101
    fmt.Println(scores)
    fmt.Println("the var i is now...", i)

    // the builtin len() will return the length
    fmt.Println("len:", len(scores))


    // declare and initialize in a single line
    highScores := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", highScores)

}

