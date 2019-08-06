package main

import "fmt"
import (
	c "github.com/egor1344/otus_home/calendar/pkg/calendar"
)

func main() {
	err := c.InitLogger()
	if err != nil {
		fmt.Println("Failed initial logger")
		return
	}
	c.ReadConfigFile("config", "./config")
}
