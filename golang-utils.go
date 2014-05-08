package goutils

import (
	"bufio"
	"bytes"
	"code.google.com/p/go-uuid/uuid"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/spaolacci/murmur3"
	"io/ioutil"
	"log"
	"net/http"
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
	hasher.Write([]byte(uuid.NewUUID()))
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

type p struct {
	a, b int
}

// SessionSignature generates a hashed sigature
// It also converts from struct to []byte
func SessionSignature(data *p) ([]byte, error) {
	h := sha256.New()
	b, err := json.Marshal(data)
	if err != nil {
		return nil, nil
	}
	_, err = h.Write(b)
	if err != nil {
		return nil, nil
	}
	return h.Sum([]byte{}), nil
}

// Make a http GET function
func HttpGet() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

// HTTP handler function
func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultClient.Do(r)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	for _, c := range resp.SetCookie {
		w.Header().Add("Set-Cookie", c.Raw)
	}
	w.WriteHeader(resp.StatusCode)
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != os.EOF {
		panic(err)
	}
	w.Write(result)
}

// Define route and write back response
func main() {
	http.HandleFunc("/", handler)
	log.Println("Start serving on port 8888")
	http.ListenAndServe(":8888", nil)
	os.Exit(0)
}

// Find any available address
func listenTCP() (net.Listener, string) {
	l, e := net.Listen("tcp", "127.0.0.1:0") // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	return l, l.Addr().String()
}

// genCode generates n digits of access code
func genCode(n uint32) string {
	rand.Seed(time.Now().Unix())
	var str string
	for i := uint32(0); i < n; i++ {
		str = fmt.Sprintf("%s%s", str, strconv.Itoa(rand.Intn(10)))
	}
	return str
}
