package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gologger/eslogger"
	"github.com/gologger/utils"
	"github.com/gologger/watcher"
)

func main() {
	utils.ClearScreen()
	// Define command options
	watcherPath := flag.String("watcherPath", "", "Define the watcher path to tell the logger where to look for logs")

	// Parse the command-line flags
	flag.Parse()

	// Retrieve command arguments
	// args := flag.Args()

	if flag.NFlag() == 0 && flag.NArg() == 0 {
		startQuestionnaire()
	} else {
		Instantiate(*watcherPath)
	}
}

func startQuestionnaire() {
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	color.New(color.FgCyan).Println("Welcome to GoLogger!")
	color.New(color.FgHiMagenta).Println("GoLogger is a tool that reads through your log file and save them in elastic search database and create visual reports based on the logs using Grafana.")
	fmt.Print("Enter the Watcher Path: (leave empty for the current directory) ")
	// Read the user input
	scanner.Scan()
	watcherPath := scanner.Text()

	Instantiate(watcherPath)
}

func Instantiate(watcherPath string) {
	// Make this one dynamic as well along with a
	// validation that tests the given constraints on realtime
	_, err := eslogger.MakeConnection()
	if err != nil {
		color.New(color.FgRed).Println(err.Error())
		os.Exit(1)
	}

	if watcherPath == "" {
		watcherPath = "./"
	}

	watcher.StartWatcher(watcherPath)
}

/*
	// I found this set of code useful because, I initially went with an option to not allow user to move further if they didn't provide args and options
	// but Now, I changed it, If the user doesn't provide args and options the app will start a questionnaire

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Usage()
	os.Exit(1)
*/
