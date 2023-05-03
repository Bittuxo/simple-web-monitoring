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

const monitor = 3
const delay = 5

func main() {
	readFile()

	for {
		intro()
		command := read()

		switch command {
		case 1:
			startScan()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Wrong Command")
			os.Exit(-1)
		}

		fmt.Println("")
	}
}

func intro() {
	fmt.Println("1 - Start Scanning")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit Software")
}

func read() int {
	readcommand := 0
	fmt.Scan(&readcommand)

	return readcommand
}

func startScan() {
	fmt.Println("Starting Scanning")

	sites := readFile()

	for i := 0; i < monitor; i++ {
		for i, site := range sites {
			fmt.Println(i, site)
			testSite(site)
			fmt.Println("")
		}
		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Status: OK")
		writeFile(site, true)
	} else {
		fmt.Println("Site:", site, "Error:", resp.StatusCode)
		writeFile(site, false)
	}
}

func readFile() []string {

	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func writeFile(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - status: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(file))
}
