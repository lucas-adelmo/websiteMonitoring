package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	
	introduction()

	for {
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
}

func introduction() {
	var version int = 1.0 // It is static typing.

	fmt.Println("Hello there! This program is in version", version)
}

func menu() {
	fmt.Println("Choose a command:")
	fmt.Println("1 - To init monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit program")
}

// The return statemant type have to be declared:
func inputCommand() int {
	var command int
	fmt.Scanf("%d", &command) // "%d" is a format specifiers and "&" is a pointer to the variable command
	return command
}

func startMonitoring() {
	var delay string
	var tests int

	fmt.Println("Type how many times you want to monitor/request the status code HTTP of each website (e.g., 1, 2, 3, ...):")
	fmt.Scan(&tests)
	fmt.Println("Type the interval of verifications (e.g., 5s, 2m, 1h):")
	fmt.Scan(&delay)
	fmt.Println("Monitoring")

	duration, err := time.ParseDuration(delay)
	if err != nil {
		fmt.Println("Invalid duration format:", err)
		return
	}

	sites := readSitesFile()

	monitoring(sites, tests, duration)

}

func monitoring(sites []string, tests int ,duration time.Duration) {
	
	for n := 1; n <= tests; n++ {
		fmt.Println("Starting the verification", n, "...")
		for i, value := range sites {
			resp, err := http.Get(sites[i])

			if err != nil {
				fmt.Println("Something went wrong here:", err)
			}

			if resp.StatusCode == 200 {
				fmt.Println(value, "is on ar. Status code: 200")
				recordingLogs(sites[i], true)
			} else {
				fmt.Println(value, "is not answering. Possible status code: 201,500-504")
				recordingLogs(sites[i], false)
			}
		}
		time.Sleep(duration)
		fmt.Println("")
	}

}

func showLogs(){
	fmt.Println("Showing logs")

	file, err := ioutil.ReadFile("log.txt")
	errHandler (err)

	fmt.Println(string(file))
}

func recordingLogs(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	errHandler (err)

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online:" + strconv.FormatBool(status) + "\n")
	file.Close()
}

func exitApp() {
	fmt.Println("Exiting program")
	os.Exit(1)
}

func commandUnknown() {
	fmt.Println("I don't reconize this command")
	os.Exit(-1)
}

func readSitesFile() []string {

	var sites []string

	// file, err := ioutil.ReadFile("sites.txt") // Show the content inside the file as bytecode (you can convert in string)

	file, err := os.Open("sites.txt") // Just show a pointer to the file
	errHandler (err)

	reader := bufio.NewReader(file) // creates a reader that roam the content inside file

	for {
		line, err := reader.ReadString('\n') // read string until line break "\n"
		line = strings.TrimSpace(line)       // ignore line break
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func errHandler (err error) {
	if err != nil {
		fmt.Println("Something went wrong here:", err)
	}
}

