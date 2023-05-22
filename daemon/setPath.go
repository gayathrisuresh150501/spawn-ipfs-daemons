package daemon

import (
	"os"
)

func SetPath(path string) error {
	return os.Setenv("IPFS_PATH", path)
	//  ExecCmd("ipfs", "daemon")
	// if err != nil {
	// 	// fmt.Println("Error setting environment variable:", err)
	// 	return err
	// }
	// return errors.New("Unable to set ENVIRONMENT VARIABLE")
}
