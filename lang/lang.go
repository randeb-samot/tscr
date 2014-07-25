package lang

import "fmt"
// define language attributes as a struct

type Lang struct {
	Name string
	Runcmd string
	Buildcmd string
}

func (l Lang) Compile() string {
	return l.Buildcmd
}

func (l Lang) Run() string {
	return l.Runcmd
}


type Error struct {
	Name string
	Err string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Err) 
}
