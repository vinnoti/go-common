package rpc

import (
	"bytes"
	"github.com/gorilla/rpc/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CallRPC(url string, methodName string, args interface{}, reply interface{}) error {
	message, err := json.EncodeClientRequest(methodName, args)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		log.Errorf(err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error in sending request to %s. %s", url, err)
		return err
	}

	err = json.DecodeClientResponse(resp.Body, &reply)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	return nil
}
