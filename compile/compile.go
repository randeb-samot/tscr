package main

import	"github.com/tscr/lang"
import	"log"
import	"fmt"
import	"os/exec"

func main() {
	lang := lang.Lang{"Java","java","javac"}

	cmd := exec.Command(lang.Compile(), "TSCR.java")
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

