package main
import "fmt"
func main() {
var a int = 200
var b int = 403
switch a {
		case 10: 
			fmt.Println("a is 10")
		case 100:
			fmt.Println("a is between 50 - 120")
		default: 
			fmt.Println("a is neither 0 nor 120")
	}
switch {
	case a + b == 233:
		fmt.Println("a plus b is equal to 233")
	case a < b:
		fmt.Println("a is less than b")
	default:
		fmt.Println("a + b is 503,and a is less than b")
	}
}
