package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

var global1 int             // Declaring a variable with `var` keyword
var global2 int = 10        // Assigning a value to a variable right away
var global3 = "this is ok!" // Go infer type based on your value
// badVar := "this doesn't work"   // `:=` can be used only inside a function

const delay = 1
const tests int = 5

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
	name := "Lucas"       // It means assign the string Lucas to the var name. It is dynamic typing.
	var version int = 1.0 // It is static typing.

	fmt.Println("Hello", name)
	fmt.Println("The program is in version", version)
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
	fmt.Println("Monitoring")

	// Array declaration:
	// var x [5] int
	// x := [5]int{2, 3, 5, 7, 11, 13}
	// Slice declaration:
	// var s [] int
	// s := []int{2, 3, 4, 9, 1}

	// sites := []string{
	// 	"https://httpstat.us/Random/200,201,500-504",
	// 	"https://www.alura.com.br/",
	// 	"https://medium.com/creators",
	// }

	sites := readSitesFile()

	monitoring(sites)

}

func monitoring(sites []string) {
	// Two way to write for conditions:
	// 1 -> i, value := range sites
	// 2 -> i := 0; i<len(sites) ; i++
	// where sites[i] = value

	for n := 0; n <= tests; n++ {
		fmt.Println("Starting the verification", n, "...")
		for i, value := range sites {
			resp, err := http.Get(sites[i])

			if err != nil {
				fmt.Println("Something went wrong here:", err)
			}

			if resp.StatusCode == 200 {
				fmt.Println("Success! The site", value, " is on ar. Status code 200")
			} else {
				fmt.Println("The site", value, " is not answering. Possible status code: 201,500-504")
			}
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
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

func readSitesFile() []string {

	var sites []string

	// file, err := ioutil.ReadFile("sites.txt") // Show the content inside the file as bytecode (you can convert in string)

	file, err := os.Open("sites.txt") // Just show a pointer to the file

	if err != nil {
		fmt.Println("Something went wrong here:", err)
	}
	reader := bufio.NewReader(file) // creates a reader that roam the content inside file
	line, err := reader.ReadString('\n') // read string until line break "\n"

	if err != nil {
		fmt.Println("Something went wrong here:", err)
	}

	fmt.Println(line)

	return sites
}
