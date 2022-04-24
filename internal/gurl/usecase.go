package gurl

import (
	"fmt"
	"os"
	"strings"
)

type Usecase struct {
	repository RepositoryImpl
}

func NewUsecase(repository RepositoryImpl) UseCaseImpl {
	return Usecase{repository: repository}
}

func UseCaseFactory() UseCaseImpl {
	return NewUsecase(RepositoryFactory())
}

// UseCaseImpl -> Interface
type UseCaseImpl interface {
	Store(url string) error
}

func (uc Usecase) Store(url string) error {
	surl := strings.Split(url, "/")
	fileName := surl[len(surl)-1]
	fmt.Println("Downloading from: ", url)

	storeFile, err := os.Create("./store/" + fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := storeFile.Close(); err != nil {
			panic(err)
		}
	}()

	return uc.repository.Store(url, storeFile)
}
