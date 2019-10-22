package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(out []byte) (int, error) {
	buffer := make([]byte, len(out))
	n, e := reader.r.Read(buffer)
	for idx := 0; idx < n; idx++ {
		if 'a' <= buffer[idx] && buffer[idx] <= 'z' {
			out[idx] = 'a' + (buffer[idx]-'a'+13)%26
		} else if 'A' <= buffer[idx] && buffer[idx] <= 'Z' {
			out[idx] = 'A' + (buffer[idx]-'A'+13)%26
		} else {
			out[idx] = buffer[idx]
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
