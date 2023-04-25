package main

import (
	"fmt"
	"os"

	"github.com/waltervargas/remindme"
)

func main() {
	if len(os.Args) > 1 {
		remindme.Add(os.Args[1])
	}
	
	rs, err := remindme.List()
	if err != nil {
		panic(err)
	}
	for _, r := range rs {
		fmt.Printf("%s\n", r.What)
	}
}
