package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func unmarshal(body io.Reader, v interface{}) error {
	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		log.Infof("error[ioutil.ReadAll]: %s\n", err)
		return err
	}
	bodyBytes = bytes.TrimPrefix(bodyBytes, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(bodyBytes, &v)

	if err != nil {
		log.Infof("error[json.Unmarshal]: %s\n", err)
		return err
	}
	return nil
}
