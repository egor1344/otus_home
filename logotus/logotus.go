package logotus

import (
	"io"
	"strconv"
	"strings"
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
	var b strings.Builder
	b.WriteString(t.Format("2019-01-01"))
	b.WriteString(" " + strconv.Itoa(ha.Id))
	b.WriteString(" " + strconv.Itoa(ha.Grade))
	b.WriteString("\n")
	return b.String()
}

type HwSubmitted struct {
	ID      int
	Code    string
	Comment string
}

func (hs *HwSubmitted) String() string {
	var t time.Time
	t = time.Now()
	var b strings.Builder
	b.WriteString(t.Format("2019-01-01"))
	b.WriteString(" " + strconv.Itoa(hs.ID))
	b.WriteString(" " + hs.Code)
	b.WriteString(" " + hs.Comment)
	b.WriteString("\n")
	return b.String()
}

func LogOtusEvent(e OtusEvent, w io.Writer) {
	io.WriteString(w, e.String())
}
