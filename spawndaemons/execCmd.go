package spawndaemons

import (
	"log"
	"os"
	"os/exec"
)

func ExecCmd(ipfs, ipfsCmd string) {
	cmd := exec.Command(ipfs, ipfsCmd)

	// Set command output to os.Stdout and os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
