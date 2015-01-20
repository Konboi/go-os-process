package main

import (
	"fmt"
	"os"
)

func main() {
	ps, err := os.FindProcess(99360)
	if err != nil {
		fmt.Printf("error is %s :", err.Error())
	}
	fmt.Print(ps.Pid)
}
