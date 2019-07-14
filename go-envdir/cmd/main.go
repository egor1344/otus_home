package main

import (
	"fmt"
	"os"
	go_envdir "github.com/egor1344/otus_home/go-envdir"
)

func main() {
	args_programm := os.Args[1:]
	fmt.Println(args_programm)
	go_envdir.GoEnvDir(args_programm[1], args_programm[2])
}
