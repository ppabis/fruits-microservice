package mocks

import "net/http"

func NewRequestWithHeaders(method string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, "http://localhost:8081", nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	return req, nil
}
