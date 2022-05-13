
// classic counting for loop
// init             condition       post statement
for ourclass := 0; ourclass < 20; ourclass++ {
     fmt.Println(ourlcass)
}

// while loop
fruit := 5
for ; fruit < 55 ; {
    fruit += 10
    fmt.Println(fruit)
}

// infinite loop
alta := 0
for {
    // stuff to do in your loop
    fmt.Println("hello!")
    // other stuff
    // conditional on when to break
    if alta == 10 {
        break
    }
}

// range (for each) loops
dataset := []{"bart", "lisa", "maggie"}
// where i is the index and v is the value
for i, v := range dataset {
    fmt.Prinln("the index is", i)
    fmt.Println("the value is", v)
}


