package main

import (
	"fmt"
	"regexp"

	"curl/internal/gurl"
)

var notValidUrlErr = "Its not a valid URL:\n\t Valid examples: https://www.google.com \t http://www.example.com/ \n Value: %s"

type Handler interface {
	Store(url string) error
}

func HandlerFactory() Handler {
	return NewHandler(gurl.UseCaseFactory())
}

func NewHandler(uc gurl.UseCaseImpl) Handler {
	return &handler{
		uc: uc,
	}
}

type handler struct {
	uc gurl.UseCaseImpl
}

func (h *handler) Store(url string) error {
	if !isUrl(url) {
		return fmt.Errorf(notValidUrlErr, url)
	}
	return h.uc.Store(url)
}

func isUrl(url string) bool {
	regex := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z\d()!@:%_\+.~#?&\/\/=]*)`)
	return regex.MatchString(url)
}
