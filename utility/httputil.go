package utility

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Request struct {
	Url     url.URL
	Headers map[string]string
	APIKey  string
}

func (r *Request) Post(body any) (*http.Response, error) {
	if r.APIKey != "" {
		r.Headers["Authorization"] = "Bearer " + r.APIKey
	}

	json, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, r.Url.String(), bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

type Response struct {
	StatusCode int
	Body       []byte
}

func (r *Response) GetBody() []byte {
	return r.Body
}

func (r *Response) GetStatusCode() int {
	return r.StatusCode
}
