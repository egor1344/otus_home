package gocopy

import (
	"fmt"
	"log"
	"bufio"
	"os"
)

func GoCopy(pathFrom string, pathTo string, limit int, offset int) error {
	fmt.Println(pathFrom, pathTo, limit, offset)
	fromfile, err := openfile(&pathFrom)
	if err != nil {
		return err
	}
	reader := bufio.NewReaderSize(fromfile, limit)
	tofile, err := createfile(&pathTo)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(tofile)
	readerwriter := bufio.NewReadWriter(reader, writer)
	log.Println(readerwriter)
	return nil
}

func openfile(pathFrom *string) (file *os.File,err error) {
	file, err = os.Open(*pathFrom)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func createfile(pathto *string) (file *os.File,err error) {
	file, err = os.Create(*pathto)
	if err != nil {
		return nil, err
	}
	return file, nil
}