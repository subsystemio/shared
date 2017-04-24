package Shared

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func ToJSON(i interface{}) *bytes.Buffer {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(i)

	return b
}

func ReadBody(body io.ReadCloser) []byte {
	s, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatalf("Error reading request body")
	}

	return s
}
