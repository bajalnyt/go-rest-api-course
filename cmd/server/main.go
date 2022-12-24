package main

import "fmt"

// Run instantiates and starts the app
func Run() error {
	fmt.Println("Starting up application")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
