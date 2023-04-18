package auth

import (
	"testing"

	"fruits_microservice/test/mocks"
)

func Test_NoTokenGiven(t *testing.T) {

	// Mock HTTP request
	mockRequest, err := mocks.NewRequestWithHeaders("GET", map[string]string{})
	if err != nil {
		t.Fatal(err)
	}

	mockResponseWriter := &mocks.MockResponseWriter{}

	token := Authenticate(mockResponseWriter, mockRequest)
	if token != nil {
		t.Fatal("Expected nil token")
	}

	if mockResponseWriter.Status != 401 {
		t.Fatal("Expected status code 401")
	}

}

func Test_GoodTokenGiven(t *testing.T) {
	privateKey := mocks.MockKeyPair()
	mockToken, err := mocks.MockJWT(1, "fruits").SignedString(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	mockRequest, err := mocks.NewRequestWithHeaders("GET", map[string]string{
		"X-Auth-Token": mockToken,
	})

	if err != nil {
		t.Fatal(err)
	}

	mockResponseWriter := &mocks.MockResponseWriter{}

	token := Authenticate(mockResponseWriter, mockRequest)

	if token == nil {
		t.Fatal("Expected token")
	}

	if mockResponseWriter.Status != 0 {
		t.Fatal("Expected status code to not be set")
	}
}

func Test_WrongService(t *testing.T) {
	privateKey := mocks.MockKeyPair()
	mockToken, err := mocks.MockJWT(1, "wrong").SignedString(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	mockRequest, err := mocks.NewRequestWithHeaders("GET", map[string]string{
		"X-Auth-Token": mockToken,
	})

	if err != nil {
		t.Fatal(err)
	}

	mockResponseWriter := &mocks.MockResponseWriter{}

	token := Authenticate(mockResponseWriter, mockRequest)

	if token != nil {
		t.Fatal("Expected nil token")
	}

	if mockResponseWriter.Status != 401 {
		t.Fatal("Expected status code 401")
	}
}

func Test_WrongSignature(t *testing.T) {
	privateKey := mocks.MockKeyPair() // First key
	mocks.MockKeyPair()               // Force new public key

	mockToken, err := mocks.MockJWT(1, "fruits").SignedString(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	mockRequest, err := mocks.NewRequestWithHeaders("GET", map[string]string{
		"X-Auth-Token": mockToken,
	})

	if err != nil {
		t.Fatal(err)
	}

	mockResponseWriter := &mocks.MockResponseWriter{}

	token := Authenticate(mockResponseWriter, mockRequest)

	if token != nil {
		t.Fatal("Expected nil token")
	}

	if mockResponseWriter.Status != 401 {
		t.Fatal("Expected status code 401")
	}

}

func Test_NoUserGiven(t *testing.T) {
	privateKey := mocks.MockKeyPair()
	mockToken, err := mocks.MockJWT(-1, "fruits").SignedString(privateKey)

	if err != nil {
		t.Fatal(err)
	}

	mockRequest, err := mocks.NewRequestWithHeaders("GET", map[string]string{
		"X-Auth-Token": mockToken,
	})

	if err != nil {
		t.Fatal(err)
	}

	mockResponseWriter := &mocks.MockResponseWriter{}

	token := Authenticate(mockResponseWriter, mockRequest)

	if token != nil {
		t.Fatal("Expected nil token")
	}

	if mockResponseWriter.Status != 401 {
		t.Fatal("Expected status code 401")
	}
}
