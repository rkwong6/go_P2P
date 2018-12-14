package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	_"gopkg.in/cheggaaa/pb.v1"
	"time"
	"io/ioutil"
	"math/rand"
)

// Each node has a download speed, upload speed, standby status, and down status
type node struct {
	nodeID      int  // unique identifier
	// downloadSpd int  // of file from seeded nodes (in Mb/sec)
	uploadSpd   int  // of file to end user (in Mb/sec)
	isStandby   bool // (as a seed) Node is either on standby (true) or in use (false)
	isDown      bool // Node is no longer a seed if true; if down
}

type fdata struct {
	data []byte
	runtime int
	nodeID int
}

/* Global variables */
var numNodes int = 10;
var dlRate int = 10;
var maxErrRate int = 10;

func main() {
	var nodeList []node;
	rand.Seed(time.Now().UTC().UnixNano());


	// Define usage statement
	usage := "[1] Run Simulation\n" +
		"[2] Check parameters\n" +
		"[3] Edit # of Nodes\n" +
		"[4] Edit Download Rate      (Mbps)\n" +
		"[5] Edit Error/Failure Rate (%)\n"

	for {
		fmt.Println("--------------------------------\n\nPeer-to-peer Simulation in Go")

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Usage:")
		fmt.Println(usage)
		fmt.Print("Enter command: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		switch {
		case text == "1":
			fmt.Println("Enter filepath: ")
			text, _ := reader.ReadString('\n');
			text = strings.TrimSuffix(text, "\n");
			data, err := ioutil.ReadFile(text);
			if err != nil {
				fmt.Println("Invalid filepath provided");
			}
			fmt.Println("Running simulation...")
			run_simulation(data, nodeList)
		case text == "2":
			check_params()
		case text == "3":
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
			nodeList = make([]node, numNodes);
			generateNodes(nodeList);
			fmt.Println("# of Nodes successfully updated to ", len(nodeList))
		case text == "4":
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
			fmt.Print("Please specify error/failure rate: ")
			maxErrRateStr, _ := reader.ReadString('\n')
			maxErrRateStr = strings.TrimSuffix(maxErrRateStr, "\n")
			maxErrRateVal, err := strconv.Atoi(maxErrRateStr)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			maxErrRate = maxErrRateVal
			fmt.Println("Erorr/failure rate successfully updated")
		default:
			fmt.Println("Error: Invalid input, please refer to usage statement...")
		}
	}

}

/* Handles main execution of simulation of a peer-to-peer network */
func run_simulation(data []byte, nodeList []node) {
	if(len(nodeList) == 0){
		fmt.Printf("Error: No nodes on network");
		return;
	}
	split := len(data)/len(nodeList);
	remainder := len(data) % len(nodeList);
	var nodeData []byte;

	nodeChan := make(chan fdata, len(nodeList));
	allData := make([]fdata, len(nodeList));


	// bar := pb.StartNew(len(nodeList))
	for i := range nodeList {
		// bar.Increment()
		if(i != 0) {
			nodeData = data[0+i*split:split+i*split];
		} else {
			nodeData = data[0+i*split:split+i*split + remainder];
		}
		go uploadFile(nodeList[i], nodeData, nodeChan);
		// time.Sleep(time.Millisecond)
		//INSERT SIMULATION HERE
	}

	for i := range allData {
		allData[i] = <- nodeChan;
	}
	for i := range allData {
		fmt.Println("Time to finish uploading", len(allData[i].data) , "bytes on node ", allData[i].nodeID ," : ", allData[i].runtime);
	}
	// bar.FinishPrint("The End!")
	// fmt.Println(split, "    ", remainder);
}

func uploadFile(node_n node, nodeData []byte, nodeChan chan<- fdata ) {
	var lagTime = rand.Intn(maxErrRate);												//amount of intitial waiting prior to beginning upload
	var uploadSpeed = rand.Intn(node_n.uploadSpd/2) + rand.Intn(node_n.uploadSpd/2);	//buffer for how much data to upload per local time unit
	var local_time = 0;

	var dataUploaded = 0;
	var lower_offset = 0;
	var upper_offset = len(nodeData);

	if(len(nodeData) < uploadSpeed) {
		upper_offset = len(nodeData);
	} else {
		upper_offset = uploadSpeed;
	}

	for dataUploaded <= len(nodeData){
		if local_time > lagTime {
			for i := lower_offset; i < upper_offset; i++ {
				dataUploaded += 1;
			}
			upper_offset += uploadSpeed;
			lower_offset += uploadSpeed;
		} 
		local_time += 1;

		time.Sleep(time.Millisecond);
	}
	var myData = fdata{nodeData, local_time, node_n.nodeID};

	nodeChan <- myData;

}

func generateNodes(nodeList []node) {
	for i := 0; i < numNodes; i++ {
		nodeList[i] = node{i, dlRate, false, false};
	}
}

/* Prints the current parameters which are set to be used in execution of simulation */
func check_params() {
	fmt.Println("---------------------------------")
	fmt.Println("Current Parameters For Simulation")
	fmt.Print("# of Nodes:    ", numNodes, "\n")
	fmt.Print("Download Rate: ", dlRate, "Mb/s\n")
	fmt.Print("Error Rate:    ", maxErrRate, "%\n")
}
