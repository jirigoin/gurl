package gurl

import (
	"fmt"
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
	sUrl := strings.Split(url, "/")
	fileName := sUrl[len(sUrl)-1]
	fmt.Println("Downloading from: ", url)
	return uc.repository.Store(url, fileName)
}
