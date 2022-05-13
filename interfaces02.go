/* RZFeeser | Alta3 Research
   Interfaces - Getting at interfaces */

package main

import "fmt"


// interfaces
type animal interface {
    breathe()
    walk()
}


// structures
type tiger struct {
     age int
}

type giraffe struct {
     age int
}


// methods - tiger

func (l tiger) breathe() {
    fmt.Println("tiger breathes")
}

func (l tiger) walk() {
    fmt.Println("tiger walk")
}

// methods - giraffe

func (l giraffe) breathe() {
    fmt.Println("giraffe breathes")
}

func (l giraffe) walk() {
    fmt.Println("giraffe walk")
}

// main
func main() {
    l := tiger{age: 10}
    callBreathe(l)
    callWalk(l)

    d := giraffe{age: 5}
    callBreathe(d)
    callWalk(d)
}


// functions that exect an interface
func callBreathe(a animal) {
    a.breathe()
}

func callWalk(a animal) {
    a.walk()
}

