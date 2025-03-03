package utils

import (
	"fmt"
	"project_spellbook/cmd/definitions"
	"sync"

	"github.com/t94j0/nmap"
)

func CompileNetworkMap(ipAddress string) {
	var waitGroup sync.WaitGroup
	//the channel that will hold the results
	//what will be in the result channel
	//
	asyncPortChannel := make(chan []nmap.Port, 10)
	//the result list
	scannedPorts := []nmap.Port{}

	//the list of ranges of up to 10000 ports
	ranges := [][2]int{
		{1, 1000},
		{1001, 2000},
		{2001, 3000},
		{3001, 4000},
		{4001, 5000},
		{5001, 6000},
		{6001, 7000},
		{7001, 8000},
		{8001, 9000},
		{9001, 10000},
	}

	//for each range, start a go routine to scan the network
	for _, r := range ranges {
		//add a channel to the weight group for each element in the range
		waitGroup.Add(1)
		go func(lowPort, highPort int) {
			//decrement the weight group when the go routine is done
			defer waitGroup.Done()
			//scan the network
			detectedPorts := definitions.MapNetwork(ipAddress, uint16(lowPort), uint16(highPort))
			//add the results to the results channel
			asyncPortChannel <- detectedPorts
		}(r[0], r[1]) //pass the range to the go routine
	}

	//start a go routine to wait for all the go routines to finish
	go func() {
		//wait for all the go routines to finish
		waitGroup.Wait()
		//close the results channel
		close(asyncPortChannel)
	}()

	//wait for the results to be added to the results channel
	for portList := range asyncPortChannel {
		scannedPorts = append(scannedPorts, portList...)
	}

	//print the results
	for _, port := range scannedPorts {
		fmt.Println(port.ToString())
	}
}
