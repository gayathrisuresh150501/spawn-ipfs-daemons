package daemon

import "fmt"

func ReqDaemonInst() int {
	var totalInstances int
	fmt.Print("Please enter the number of daemons to be initiated: ")
	fmt.Scanf("%d", &totalInstances)

	return totalInstances
}
