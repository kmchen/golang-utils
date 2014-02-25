package goutils

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"
)

// Uint64ToByte converts an uint64 to bytes array in BigEndian
func Uint64ToByte(data uint64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, data)
	return buf.Bytes()
}

// ByteToUint64 converts an uint64 in bytes to int64 with BigEndian
func ByteToUint64(data []byte) uint64 {
	var value uint64
	buf := bytes.NewReader(data)
	binary.Read(buf, binary.BigEndian, &value)
	return value
}

// RandomUint64 creates a random uint64 number
func RandomUint64(data string) uint64 {
	hasher := murmur3.New64()
	hasher.Write([]byte(NewUUID()))
	return hasher.Sum64()
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
