/* Alta3 Research | RZFeeser
   Refer - Basic patterns - panic(), defer, and recover()  */


package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

// f figures out g made a mistake, and prevents main from crashing
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

// error happens when g is called repeatedly by f
func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}

