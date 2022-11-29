package jzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Unzip(data []byte) (resData []byte, err error) {
	bData := bytes.NewBuffer(data)

	var reader io.Reader
	reader, err = gzip.NewReader(bData)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(reader)
	if err != nil {
		return
	}

	resData = resB.Bytes()

	return
}
