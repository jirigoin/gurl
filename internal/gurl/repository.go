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
	Store(url, filename string) error
}

func NewRepository(c *http.Client) RepositoryImpl {
	return Repository{client: c}
}

func RepositoryFactory() RepositoryImpl {
	c := http.DefaultClient
	return NewRepository(c)
}

func (r Repository) Store(url, filename string) error {
	file, err := os.Create("./store/" + filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
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
	n, err := io.Copy(file, response.Body)
	if err != nil {
		log.Fatalln("error while downloading", url, "-", err)
		return err
	}
	fmt.Printf("Store %d bytes", n)
	return nil
}
