package logotus

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type OtusEvent interface {
	String() string
}

type HwAccepted struct {
	Id    int
	Grade int
}

func (ha *HwAccepted) String() string {
	var t time.Time
	t = time.Now()
	s := t.Format("2019-01-01")
	s += " " + strconv.Itoa(ha.Id)
	s += " " + strconv.Itoa(ha.Grade)
	return s
}

type HwSubmitted struct {
	ID      int
	Code    string
	Comment string
}

func (hs *HwSubmitted) String() string {
	var t time.Time
	t = time.Now()
	s := t.Format("2019-01-01")
	s += " " + strconv.Itoa(hs.ID)
	s += " " + hs.Code
	s += " " + hs.Comment
	return s
}

func LogOtusEvent(e OtusEvent, w io.Writer) {
	switch e.(type) {
	case *HwAccepted:
		fmt.Printf("%v\n", e)
	case *HwSubmitted:
		fmt.Printf("%v\n", e)

	}
}
