package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Each node has a download speed, upload speed, standby status, and down status
type node struct {
	nodeID      int  // unique identifier
	downloadSpd int  // of file from seeded nodes (in Mb/sec)
	uploadSpd   int  // of file to end user (in Mb/sec)
	isStandby   bool // (as a seed) Node is either on standby (true) or in use (false)
	isDown      bool // Node is no longer a seed if true; if down
}

/* Global variables */
var numNodes int
var dlRate int
var errRate int

func main() {

	// Define usage statement
	usage := "[1] Run Simulation\n" +
		"[2] Check parameters\n" +
		"[3] Edit # of Nodes\n" +
		"[4] Edit Download Rate      (Mbps)\n" +
		"[5] Edit Error/Failure Rate (%)\n"

	for {
		fmt.Println("\n--------------------------------\n\nPeer-to-peer Simulation in Go")

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Usage:")
		fmt.Println(usage)
		fmt.Print("Enter command: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		fmt.Println("Command Entered:", text)

		switch {
		case text == "1":
			fmt.Println("Running simulation...")
			run_simulation()
		case text == "2":
			check_params()
		case text == "3":
			fmt.Println("Editing Number of Nodes...")
			fmt.Print("Please specify # of Nodes: ")
			numNodesStr, _ := reader.ReadString('\n')
			numNodesStr = strings.TrimSuffix(numNodesStr, "\n")
			numNodesVal, err := strconv.Atoi(numNodesStr)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			numNodes = numNodesVal
			fmt.Println("# of Nodes successfully updated")
		case text == "4":
			fmt.Println("Editing download rate...")
			fmt.Print("Please specify download rate: ")
			dlRateStr, _ := reader.ReadString('\n')
			dlRateStr = strings.TrimSuffix(dlRateStr, "\n")
			dlRateVal, err := strconv.Atoi(dlRateStr)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			dlRate = dlRateVal
			fmt.Println("download rate successfully updated")
		case text == "5":
			fmt.Println("Editing error/failure rate...")
			fmt.Print("Please specify error/failure rate: ")
			errRateStr, _ := reader.ReadString('\n')
			errRateStr = strings.TrimSuffix(errRateStr, "\n")
			errRateVal, err := strconv.Atoi(errRateStr)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			errRate = errRateVal
			fmt.Println("Erorr/failure rate successfully updated")
		default:
			fmt.Println("Error: Invalid input, please refer to usage statement...")
		}
	}

}

/* Handles main execution of simulation of a peer-to-peer network */
func run_simulation() {

}

/* Prints the current parameters which are set to be used in execution of simulation */
func check_params() {

}
