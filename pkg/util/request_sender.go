package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	GET_METHOD    = "GET"
	POST_METHOD   = "POST"
	PUT_METHOD    = "PUT"
	DELETE_METHOD = "DELETE"

	JSON_CONTENT_TYPE = "application/json"
)

type RequestSender struct {
	Url         string
	Method      string
	ContentType string
	Content     []byte
}

func (r *RequestSender) Send() (*http.Response, map[string]interface{}, error) {
	req, err := http.NewRequest(r.Method, r.Url, bytes.NewBuffer(r.Content))
	req.Header.Set("Content-Type", r.ContentType)
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return resp, result, nil
}
