package main

import (
	"bufio"
	"fmt"
	"ipfs-n-daemon-pkg/spawndaemons"
	"os"
	"sync"
)

func main() {
	nodeCount := spawndaemons.ReqDaemonInst()
	spawndaemons.CreateNewInstance(nodeCount)

	var wg sync.WaitGroup

	for _, path := range spawndaemons.InstWithPaths {
		wg.Add(1)
		// fmt.Println("IPFS PATH :", "C:"+path)
		if err := spawndaemons.SetPath("C:" + path); err != nil {
			fmt.Println("SetPath Err:", err)
		}
		spawndaemons.ExecuteCommand(&wg)
	}

	wg.Wait()

	fmt.Println("Press ctrl + c to exit.")
	spawndaemons.ShutDaemon(spawndaemons.InstWithPaths)

	//use keyboard package

	// var termSig string
	// spawndaemons.ExecCmd("ipfs", "shutdown")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	//use scanf

	// spawndaemons.ReqDaemonInst()
	// spawndaemons.ExecCmd("ipfs", "daemon")
}

/*
keyboard.Listen(func(key keys.Key) (stop bool, err error) {
	if key.Code == keys.CtrlC {
		return true, nil // Stop listener by returning true on Ctrl+C
	}

	fmt.Println("\r" + key.String()) // Print every key press
	return false, nil // Return false to continue listening
})
*/
