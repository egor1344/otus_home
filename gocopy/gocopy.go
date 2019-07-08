package gocopy

import (
	"flag"
	"fmt"
)

var pathFrom string
var pathTo string
var limit int
var offset int

func init() {
	flag.StringVar(&pathFrom, "from", "", "path from copy file")
	flag.StringVar(&pathTo, "to", "", "path where copy file")
	flag.IntVar(&limit, "limit", 0, "Limit bytes for copy file")
	flag.IntVar(&offset, "offset", 0, "offset bytes for copy file")
}

func GoCopy(pathFrom string, pathTo string, limit int, offset int) {
	flag.Parse()
	fmt.Println(pathFrom, pathTo, limit, offset)
}
