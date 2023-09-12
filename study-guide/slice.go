package main 
import "fmt"
func main() {		

	arr := [10]int{10,20,30,40,50,60,70,80,90,100}
	slice_1 := arr[0:8]
	fmt.Println(slice_1)
	fmt.Println(cap(slice_1))
	 fmt.Println(len(slice_1))
	sub_slice := slice_1[0:3]
	fmt.Println(sub_slice)

	slice_1 = append (slice_1, 500, 700)	
	 fmt.Println("after appending")
	 fmt.Println(cap(slice_1))
	  fmt.Println(len(slice_1))
	   fmt.Println(slice_1)

	arr_2 := [5]int{1,3,5,7,9}
	slice_0 := arr_2[0:2]
	 fmt.Println(slice_0)
	 fmt.Println(arr_2)
	 
}
