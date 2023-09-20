package main

import (
    "fmt"
    "math/rand"
    "time"
    "math"
    
)

func main() {
    // Seed the random number generator with the current time
    rand.Seed(time.Now().UnixNano())

    // Generate a random integer between 1 and 100
    randomNumber := rand.Intn(100) + 1
    fmt.Printf("the square of 9: %g\n", math.Sqrt(9))
    fmt.Println("Random number between 1 and 100:", randomNumber)
}
