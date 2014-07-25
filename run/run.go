package main

import "os/exec"
import 	"fmt"
import 	"log"

// add rpc server to run commands according to requests


func main() {
	cmd := exec.Command("java", "TSCR")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	buff := make([]byte,1000)
	stdout.Read(buff)	
	fmt.Println( string(buff))
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

