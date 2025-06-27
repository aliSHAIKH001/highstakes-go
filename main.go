package main

import "fmt"

func getName() string {
	name := "Ali Shaikh"

	fmt.Println("Welcome to Ali's Casino...")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s lets play\n", name)
	return name
}

func main() {
	getName()
}