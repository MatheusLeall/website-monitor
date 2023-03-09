package main

import (
	"fmt"
	"os"
)

func main() {
	introduction()
	showCommandOptions()

	userCommand := readUserCommand()

	switch userCommand {
	case 1:
		fmt.Println("Starting monitor...")
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exiting...")
		exitSafely()
	default:
		fmt.Println("Invalid option, exiting...")
		exitSafely()
	}
}

func introduction() {
	var name string
	var systemVersion float32 = 1.1

	fmt.Printf("Let's start, what is your name: ")
	fmt.Scan(&name)

	fmt.Println("\nHi", name)
	fmt.Println("This system is in version", systemVersion)
}

func showCommandOptions() {
	fmt.Println("\n1- Start monitor")
	fmt.Println("2- Show monitor logs")
	fmt.Println("0- Exit")
}

func readUserCommand() int {
	var userCommand int8
	fmt.Scan(&userCommand)

	return int(userCommand)
}

func exitSafely() {
	os.Exit(0)
}
