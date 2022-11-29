package jzip

import (
	"bytes"
	"compress/gzip"
)

func Zip(data []byte) (compressedData []byte, err error) {
	var buffer bytes.Buffer
	gz := gzip.NewWriter(&buffer)

	_, err = gz.Write(data)
	if err != nil {
		return
	}

	err = gz.Flush()
	if err != nil {
		return
	}

	err = gz.Close()
	if err != nil {
		return
	}

	compressedData = buffer.Bytes()
	return
}
