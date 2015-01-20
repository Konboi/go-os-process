package main

import (
	"fmt"
	"os"
)

func main() {
	ps, err := os.FindProcess(54327)

	if err != nil {
		fmt.Printf("error is %s :", err.Error())
	}
	fmt.Print(ps.Pid)

	err = ps.Kill()
	if err != nil {
		fmt.Printf("kill process error is %s :", err.Error())
	}
}
