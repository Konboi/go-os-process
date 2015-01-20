package main

import (
	"fmt"
	"os"
)

func main() {
	var proc_attr os.ProcAttr

	ps, err := os.StartProcess("golang-process", ["some sttring"], proc_attr)
	)
	if err != nil {
		fmt.Printf("error is %s :", err.Error())
	}
	fmt.Print(ps.Pid)
}
