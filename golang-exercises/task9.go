package main
import (
    "fmt"
    "strconv"
        )

func main() {
    profile := map[string]string{
    "first Name":"marvellous",
    "Last Name" : "Ezemba",
    "age": strconv.FormatUint(uint64(24), 10),
    }
    fmt.Print(profile)
}