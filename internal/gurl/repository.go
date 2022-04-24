package gurl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Repository struct {
	client *http.Client
}

type RepositoryImpl interface {
	Store(url string, storeFile *os.File) error
}

func NewRepository(c *http.Client) RepositoryImpl {
	return Repository{client: c}
}

func RepositoryFactory() RepositoryImpl {
	c := http.DefaultClient
	return NewRepository(c)
}

func (r Repository) Store(url string, storeFile *os.File) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("bad status code: %s", response.Status)
	}
	n, err := io.Copy(storeFile, response.Body)
	if err != nil {
		log.Fatalln("error while downloading", url, "-", err)
		return err
	}
	fmt.Printf("Store %d bytes", n)
	return nil
}
