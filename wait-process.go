package main

import (
	"fmt"
	"os"
)

func main() {
	ps, err := os.FindProcess(51940)

	if err != nil {
		fmt.Printf("error is %s :", err.Error())
	}

	status, err := ps.Wait()
	if err != nil {
		fmt.Printf("kill process error is %s :", err.Error())
	}

	fmt.Printf("status is: %s \n", status.String())

}
