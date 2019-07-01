package parallel

import (
	"testing"
	"errors"
	"fmt"
	"log"
)

func TestParralel(t *testing.T) {
	w1 := func () error {
		fmt.Println("f1")
		return nil
	}
	w2 := func() error {
		fmt.Println("f2")
		return nil
	}
	w3 := func() error {
		fmt.Println("f3")
		return nil
	}
	errorFunc := func() error {
		fmt.Println("Error func")
		return errors.New("Error func")
	}

	err := Parallel(1, 0, []CustomFunc{w1, w2, w3})
	if err != nil {log.Println(err)}
	log.Println("1------------------------------")
	err = Parallel(2, 2, []CustomFunc{w1, w2, w3, w1, errorFunc, w3, w1, errorFunc, w3})
	if err != nil {log.Println(err)}
	log.Println("2------------------------------")
	err = Parallel(3, 1, []CustomFunc{w1, w2, errorFunc, w1, errorFunc, w3, w1, errorFunc, w3})
	log.Println("3------------------------------")
	if err != nil {log.Println(err)}
}