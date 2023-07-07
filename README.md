<h1 align="center"> Website Monitoring </h1>

<p align="center">
  <img src="./websiteMonitoring.png" />
</p>

# Introduction

This simple CLI (Command-Line Interface) program monitors the HTTP status code of a website. It's my first project using Go programming language. It's part of a course from [Alura](https://www.alura.com.br/).

# How it works?

### - **First step** 

Paste each URL you want to monitor inside the file **sites.txt**. Each line of the file should contain a URL.

*sites.txt*
```txt
https://cursos.alura.com.br/dashboard
https://httpstat.us/Random/200,201,500-504
https://httpstat.us/200?sleep=5000
```
### - **Second step** 

Run the program using the command bellow in the terminal:

```bash
go run index.go
```
It presents a menu with three options:
``` bash
Hello there! This program is in version 1
Choose a command:
1 - To init monitoring
2 - Show logs
3 - Exit program
```
If you type the digit "**1**" in the terminal, the program will ask you to choose some parameters. It's self-explanatory:
```
Type how many times you want to monitor/request the status code HTTP of each website (e.g., 1, 2, 3, ...):
```

```bash
Type the interval of verifications (e.g., 5s, 2m, 1h):
1s
```
Then, the program will execute. After the verifications you can see the result in **log.txt** and the console.To exit the program, just type the digit "**0**". 

Pretty simple right? 

# The code

As it's my first code in Go, I wanna share with you what I learned. Go main program consists roughly of three parts: package name, imported packages and main function. Like in C/C++, Go compiler always look for main() function to start from. 

The Go's standart library is pretty complete, you don't need a giant folder as node_modules. The modules that I used are: 

- "bufio": This module provides buffered I/O operations, allowing you to read and write data more efficiently by reducing the number of direct system calls. It includes types like Scanner and ReadWriter, which make it easier to work with input/output streams.

- "fmt": Is used for formatted I/O operations. It provides functions for printing and formatting text, such as Println, Printf, and Sprintf. It also includes the Scan and Scanf functions for reading formatted input.

- "io": Provides basic input/output functionality. It defines interfaces like Reader and Writer, which are implemented by various types in Go's standard library. This module is used for working with streams of data, reading from and writing to different sources or destinations.

- "io/ioutil": Provides utility functions for I/O operations. It includes functions like ReadFile and WriteFile, which allow you to read or write the contents of a file in a single function call. It simplifies common file I/O tasks by handling the details of file opening, closing, and error handling.

- "net/http": Is used for making HTTP requests and building HTTP servers. It provides functions and types for working with HTTP protocols, including sending requests, handling responses, setting headers, and managing cookies. It is commonly used for web development and interacting with web APIs.

- "os": Provides a way to interact with the operating system. It includes functions for working with files and directories, environment variables, command-line arguments, and other operating system-related functionality. It allows you to perform operations like file opening, renaming, deletion, and directory listing.

- "strconv": Provides functions for converting strings to different types and vice versa. It includes functions like Atoi (string to integer), Itoa (integer to string), ParseFloat (string to float), and FormatFloat (float to string). It is commonly used for parsing and formatting numerical values.

- "strings": Provides functions for manipulating strings. It includes functions for string concatenation, splitting, trimming, replacing, and searching. It also provides utilities for string comparisons, checking prefixes and suffixes, and converting case.

- "time": Provides functionality for working with time and durations. It includes types like Time and Duration, as well as functions for formatting and parsing time, calculating differences between time points, and performing time-related operations. It is used for tasks such as scheduling, measuring time intervals, and handling time-related calculations.

Go is a programming language known for its simplicity and elegance. In Go, you don't need to use semicolons at the end of each line, and you don't need parentheses for conditions in the **"for"** loop or **"if"** statements. The syntax of Go is loosely derived from C, but it includes additional features such as garbage collection (similar to Java), type safety, and some dynamic-typing capabilities. It is worth noting that Go does not have a **"while"** loop; instead, the **"for"** loop is used for all types of loops in Go.

I recomend [this](https://medium.com/code-zen/go-or-how-i-came-to-love-static-types-again-part-1-32120a7f599f) article to understand the basic principles of Go.

# Contact

I you want to contribute or send me Hello, just send an email.

- lucasadelmo2@gmail.com

