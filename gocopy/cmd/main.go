package main

import (
	"errors"
	"flag"
	"log"

	"github.com/egor1344/otus_home/gocopy"
)

var pathFrom string
var pathTo string
var limit int
var offset int

func init() {
	flag.StringVar(&pathFrom, "from", "", "path from copy file")
	flag.StringVar(&pathTo, "to", "", "path where copy file")
	flag.IntVar(&limit, "limit", 1024, "Limit bytes for copy file")
	flag.IntVar(&offset, "offset", 1024, "offset bytes for copy file")
}

func main() {
	flag.Parse()
	if pathFrom != "" && pathTo != "" {
		err := gocopy.GoCopy(pathFrom, pathTo, limit, offset)
		if err != nil {
			log.Println(err)
		}
	} else {
		err := errors.New("not from (and/or) to file")
		log.Println(err)
	}

}
