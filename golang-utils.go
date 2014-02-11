package goutils

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"
)

// ReadByteToInt64 converts an int64 in bytes to int64 with BigEndian
func ReadByteInt64(data []byte) (int64, error) {
	var value int64
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, err
}

// Realn reads line by line from a bufio.Reader
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

// WriteByteToFile writes an array of []byte to a file
func WriteByteToFile(filename string, data [][]byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, v := range data {
		_, err := file.Write(v)
		if err != nil {
			return err
		}
	}
	return nil
}
