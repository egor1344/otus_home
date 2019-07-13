package gocopy

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func GoCopy(pathFrom string, pathTo string, limit int, offset int) error {
	fmt.Println(pathFrom, pathTo, limit, offset)
	if offset > limit {
		return errors.New("offset great that limit")
	}
	fromfile, err := openfile(&pathFrom)
	if err != nil {
		return err
	}
	stat, err := fromfile.Stat()
	if err != nil {
		return err
	}
	sizefile := stat.Size()
	tofile, err := createfile(&pathTo)
	if err != nil {
		return err
	}
	buf := make([]byte, limit)
	offsetint64 := int64(offset)
	sizefilefloat := float32(sizefile)
	for {
		n, err := fromfile.ReadAt(buf, offsetint64)
		offsetint64 += int64(n)
		if err == io.EOF {
			_, err := tofile.Write(buf)
			if err == io.EOF {
				return err
			}
			log.Println("Запист файла окончена")
			break
		}
		if err != nil {
			return err
		}
		_, err = tofile.Write(buf)
		if err != nil {
			return err
		}
		offsetfloat := float32(offsetint64)
		percent := (offsetfloat / sizefilefloat) * 100
		log.Printf("\x0cЗаписанно %g%%", percent)
	}
	return nil
}

func openfile(pathFrom *string) (file *os.File, err error) {
	file, err = os.Open(*pathFrom)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func createfile(pathto *string) (file *os.File, err error) {
	file, err = os.Create(*pathto)
	if err != nil {
		return nil, err
	}
	return file, nil
}
