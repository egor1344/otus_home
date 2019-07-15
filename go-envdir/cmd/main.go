package main

import (
	"github.com/egor1344/otus_home/go-envdir"
	"log"
	"os"
)

func main() {
	argsProgramm := os.Args[1:]
	if len(argsProgramm) < 2 {
		log.Fatalln("Параметров меньше двух")
	}
	err := go_envdir.GoEnvDir(argsProgramm[0], argsProgramm[1])
	if err != nil {
		log.Fatalln("Параметров меньше двух", err)
		return
	}
}
