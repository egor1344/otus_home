package parallel

import (
	"testing"
	"fmt"
	"log"
)

func TestParralel(t *testing.T) {
	w1 := func () error {
		fmt.Println("First func")
		return nil
	}
	w2 := func() error {
		fmt.Println("second func")
		return nil
	}
	w3 := func() error {
		fmt.Println("three func")
		return nil
	}
	//w1()
	//w2()
	//w3()
	err := Parallel(2, 1, []CustomFunc{w1, w2, w3})
	if err != nil {log.Println(err)}
}