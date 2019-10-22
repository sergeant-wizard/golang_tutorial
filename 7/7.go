package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(out []byte) (int, error) {
	for idx := range out {
		out[idx] = 'A'
	}
	return len(out), nil
}

func main() {
	reader.Validate(MyReader{})
}
