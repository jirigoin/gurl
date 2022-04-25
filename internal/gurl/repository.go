package gurl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jirigoin/gurl/internal/rest"
)

type Repository struct {
	client rest.HTTPClient
}

type RepositoryImpl interface {
	Store(url, filename string) error
}

func NewRepository(c rest.HTTPClient) RepositoryImpl {
	return Repository{client: c}
}

func RepositoryFactory() RepositoryImpl {
	return NewRepository(rest.Client)
}

func (r Repository) Store(url, filename string) error {
	response, err := rest.Client.Get(url)
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

	_, here, _, _ := runtime.Caller(0)
	fileDir := filepath.Join(filepath.Dir(here), "../../store/", filename)
	file, err := os.Create(fileDir)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	n, err := io.Copy(file, response.Body)
	if err != nil {
		log.Fatalln("error while downloading", url, "-", err)
		return err
	}
	fmt.Printf("Store %d bytes", n)
	return err
}
