package main

import "fmt"

func main() {
	var name string
	score := 0 //named a variable to help us calculate or score as we answer questions correctly.
	num_of_ques := 3
	fmt.Printf("Welcome to my quiz game!, please type in your name: ") //print a new line after thr ptogram
	//fmt.Printf("hello %v, how are you doing? you're %v!", name, age) //%v is a place holder for name,we mainly
	//use this print format for task like this.

	fmt.Scan(&name) //we are scanning the name written and storing the value on our name variable declared.

	fmt.Printf("hello %v, welcome to the game,", name)

	var age uint
	fmt.Printf("how old are you? ")

	fmt.Scan(&age) //scans the number  written and stores it in the variable named age.

	fmt.Println(" great you are ", age, "years of age!") //returns a message with the age value.

	if age >= 10 {
		fmt.Println("you're eligible to play. Enjoy")
	} else {
		fmt.Println("Sorry you can't proceed cause you're not of age")
		return
	}
	fmt.Printf("1. what is the capital of Nigeria? ")

	var answer string
	fmt.Scan(&answer)

	if answer == "abuja" {
		fmt.Println("correct!")
		score++
	} else {
		fmt.Println("wrong!")
	}

	fmt.Printf("which is better RTX 3080 or RTX 3090: ")
	var answer2 string
	var answer2b string
	fmt.Scan(&answer2, &answer2b)
	if answer2+" "+answer2b == "RTX 3080" || answer2+" "+answer2b == "rtx 3080" {
		fmt.Println("correct!")
		score++
	} else {
		fmt.Print("wrong!")
	}
	fmt.Printf("how many vCPU does t2-micro EC2 instance for ubuntu come with: ")
	var cores uint
	fmt.Scan(&cores)
	if cores == 1 {
		fmt.Println("correct!")
		score++
	} else {
		fmt.Println("OLODO!!!!")
	}
	fmt.Printf("you score %v out of %v!", score, num_of_ques)
	percent := (float64(score) / float64(num_of_ques)) * 100
	fmt.Printf("you got %v%%\n", percent)
}
