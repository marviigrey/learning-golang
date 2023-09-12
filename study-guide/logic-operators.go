package main

import "fmt"

func main() {
 var w, d int = 3, 12
 fmt.Println((w > 2) && (w > 1))
 fmt.Println((w < 13)&& (2 > w))
 fmt.Println((w == 3) || (w != 3))
 fmt.Println((w > 4) || (w >= 5))
 fmt.Println(!(w > d))
 fmt.Println(!(true))
 }
