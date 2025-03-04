package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (mr MyReader) Read(b []byte) (int, error) {
	for i := range len(b) {
		b[i] = 'A'
	}
	return len(b), nil
}

func seven() {
	reader.Validate(MyReader{})
}
