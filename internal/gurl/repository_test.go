package gurl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/jirigoin/gurl/internal/rest"
	"github.com/stretchr/testify/assert"
)

var (
	filename   = "url_test"
	url        = "http:www.test/" + filename
	pathToFile = "../../store/" + filename
)

func mockRestClient() *rest.MockClient {
	return &rest.MockClient{}
}

func TestRepository_Store_ResponseError(t *testing.T) {
	// Given
	body := ioutil.NopCloser(bytes.NewReader([]byte("")))
	rest.GetFunc = func(s string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 400,
			Body:       body,
		}, nil
	}
	repo := NewRepository(mockRestClient())

	// When
	err := repo.Store(url, filename)

	// Then
	assert.NotNil(t, err)
}

func TestRepository_Store_Successfully(t *testing.T) {
	// Given
	t.Cleanup(deleteFile)
	rest.Client = &rest.MockClient{}
	r := "this is the response"
	b := ioutil.NopCloser(bytes.NewReader([]byte(r)))
	rest.GetFunc = func(s string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       b,
		}, nil
	}
	repo := NewRepository(mockRestClient())

	// When
	err := repo.Store(url, filename)

	// Then
	assert.Nil(t, err)
	assert.True(t, isFileCreatedSuccessfully(r))
	fmt.Println("")
}

func isFileCreatedSuccessfully(body string) bool {
	f, err := os.ReadFile(pathToFile)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	compares := body == string(f)
	return compares
}

func deleteFile() {
	err := os.Remove(pathToFile)
	if err != nil {
		log.Fatalln("Can't delete the file. Error: ", err)
	}
}
