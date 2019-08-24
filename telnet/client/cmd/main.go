package main

import (
	"errors"
	"flag"
	"log"
	"otus_home/telnet/client"
)

var addr string
var port string
var timeout int

func init() {
	flag.StringVar(&addr, "addres", "", "addres server")
	flag.StringVar(&port, "port", "", "port server")
	flag.IntVar(&timeout, "timeout", 30, "timeout context")
}

func main() {
	flag.Parse()
	if addr != "" && port != "" {
		err := client.Connect(addr, port, timeout)
		if err != nil {
			log.Println(err)
		}
	} else {
		err := errors.New("Not set addr and port")
		log.Println(err)
	}
}
