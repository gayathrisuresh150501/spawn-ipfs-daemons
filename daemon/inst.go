package daemon

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
)

const (
	DefaultPath   = `/Users/gayat/` //make it work for all systems like cd ~   ~ for %HOME%
	RootConfigDir = ".ipfs"
	RootNodePath  = `/Users/gayat/.ipfs`
)

var InstWithPaths = make(map[string]string)

func ExecuteCommand(wg *sync.WaitGroup) {
	defer wg.Done()

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("ipfs", "daemon")
	case "linux", "darwin":
		cmd = exec.Command("sh", "-c", "ipfs daemon")
	default:
		log.Fatal("Unsupported operating system")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		panic(err)
		// log.Fatal(err)
	}

	execPath := os.Getenv("IPFS_PATH")

	fmt.Printf("Command: %s\nOutput: %s\n", "ipfs daemon on", execPath)
}

func CreateNewInstance(instanceCount int) {
	os.Chdir(DefaultPath)
	if _, err := os.Stat(RootConfigDir); err == nil {
		InstWithPaths[RootConfigDir] = RootNodePath
	} else {
		os.Mkdir(RootConfigDir, 0755)

	}

	for i := 1; i < instanceCount; i++ {
		instName := ".ipfs" + strconv.Itoa(i)
		_, ok := InstWithPaths[instName]
		if !ok {
			err1 := os.Mkdir(instName, 0755)
			if err1 != nil {
				fmt.Println(err1)
			}
			setIPFSPath := DefaultPath + instName
			InstWithPaths[instName] = setIPFSPath
			SetPath(setIPFSPath) //handle error
			ExecCmd("ipfs", "init")
			AssignPort()
			// startDaemon()

		}

	}

}

// func ShutDaemon(InstanceWithPaths map[string]string) {
// 	for _, path := range InstanceWithPaths {
// 		err := os.Chdir(path)
// 		if err != nil {
// 			fmt.Println("Unable to shutdown daemon: ", err)
// 		}
// 		ExecCmd("ipfs", "shutdown")
// 	}
// }
