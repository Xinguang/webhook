package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

func unmarshal(body io.Reader, v interface{}) error {
	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		logger.Printf("error[ioutil.ReadAll]: %s\n", err)
		return err
	}
	bodyBytes = bytes.TrimPrefix(bodyBytes, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(bodyBytes, &v)

	if err != nil {
		logger.Printf("error[json.Unmarshal]: %s\n", err)
		return err
	}
	return nil
}
