package mocks

import "net/http"

type MockResponseWriter struct {
	Status int
}

func (m *MockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m *MockResponseWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.Status = statusCode
}
