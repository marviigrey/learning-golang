package main 
	import "fmt"

	func main() {
		test := map[string]string{"en": "English", "fr": "French"}
		fmt.Println(test["en"])
		new_map := map[string]string{"age": "12", "score": "63", "subjects": "4"}
		value, found := new_map["age"]
		fmt.Println(found, value)
		value ,found = test["fr"]
		fmt.Println(found, value)
		fmt.Println(new_map)
		test["it"] = "Italian"
		fmt.Println(test)
		new_map["age"] = "23"
		fmt.Println(new_map)
		delete(new_map, "subjects")
		fmt.Println(new_map)
		for key, value := range new_map {
			fmt.Println(key, "=>", value)
		} 

		for key := range new_map {
			delete(new_map, key)
		}
	}
