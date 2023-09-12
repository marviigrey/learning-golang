package main
 import "fmt"
 func main() {
	 grades := [...] int{60, 70, 80}
	 grades[2] = 120
	var transport [4] string = [4] string{"house", "cars", "plane", "boat"}
	fmt.Println(transport, grades)
	fmt.Println(len(grades))
	fmt.Println(grades[0])
	for i, element := range transport {
		fmt.Printf("%d: %d\n", i, element)
	}
	arr := [4][2] int{{12, 14}, {15, 17}, {10, 43}, {51, 32}}
	fmt.Println(arr[0][1])
	slice := []int{1,2,3,4}
	fmt.Println(slice)
 }

