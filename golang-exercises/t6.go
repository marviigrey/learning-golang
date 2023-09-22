package main
import "fmt"

func main() {
    integers := [10]int{}
    strings := [...]string{"marv", "thyron", "jp", "ebele"}
    //you can't miss different string in one data type.
    //you can add items but you can't more than 10 items to the "integers" array.
    integers[0] = 12
    games := [...]string{"Donkey Kong"}
     
    movies := [...]string{"matrix", "naruto", "superman"}
    fmt.Println(strings, integers, movies, movies[1],games, len(movies))
    
}