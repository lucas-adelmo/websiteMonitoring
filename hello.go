package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	introduction()
	menu()
	command := inputCommand()

	switch command {
	case 1:
		startMonitoring()
	case 2:
		showLogs()
	case 0:
		exitApp()
	default:
		commandUnknown()
	}
}

func introduction() {
	name := "Lucas"
	version := 1.0

	fmt.Println("Hello", name)
	fmt.Println("The program is in version", version)
}

func menu() {
	fmt.Println("Choose a command:")
	fmt.Println("1 - To init monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit program")
}

func inputCommand() int {
	var command int
	fmt.Scanf("%d", &command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitoring")
	site := "https://www.alura.com.br/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Success")
	} else {
		fmt.Println("Something went wrong")
	}
}

func showLogs() {
	fmt.Println("Showing logs")
}

func exitApp() {
	fmt.Println("Exiting program")
	os.Exit(1)
}

func commandUnknown() {
	fmt.Println("I don't reconize this command")
	os.Exit(-1)
}
