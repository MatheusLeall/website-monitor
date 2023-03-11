package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const AMOUNT_MONITORING = 5
const SECONDS_BETWEEN_REQUESTS = 3

func main() {
	introduction()
	for {
		showCommandOptions()

		userCommand := readUserCommand()

		switch userCommand {
		case 1:
			startMonitoring()
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

func startMonitoring() {
	fmt.Println("Starting monitor...")

	websites := getFileContent()

	for i := 0; i < AMOUNT_MONITORING; i++ {
		for _, siteURL := range websites {
			requestAndVerifySiteURL(siteURL)
		}
		time.Sleep(SECONDS_BETWEEN_REQUESTS * time.Second)
	}
}

func requestAndVerifySiteURL(siteURL string) {
	response, _ := http.Get(siteURL)

	if response.StatusCode == 200 {
		fmt.Println(siteURL, "is up!")
	} else {
		fmt.Println(siteURL, "is down!\nStatus:", response.StatusCode)
	}
}

func getFileContent() []string {
	var websites []string
	file, err := os.Open("resources/sites.txt")

	if err != nil {
		fmt.Println("An error occurs during file open ->", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		websites = append(websites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return websites
}
