package gurl

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUseCase_Store_Successful(t *testing.T) {
	// Given
	filename := "url_test"
	url := "http:www.test/" + filename

	controller := gomock.NewController(t)
	repositoryMock := NewMockRepositoryImpl(controller)
	repositoryMock.EXPECT().Store(url, filename).Return(nil)
	uc := NewUsecase(repositoryMock)

	// When
	err := uc.Store(url)

	// Then
	assert.Nil(t, err)
}

func TestUseCase_Store_ReturnsAnError(t *testing.T) {
	// Given
	filename := "url_test"
	url := "http:www.test/" + filename
	errMock := errors.New("mock-error")

	controller := gomock.NewController(t)
	repositoryMock := NewMockRepositoryImpl(controller)
	repositoryMock.EXPECT().Store(url, filename).Return(errMock)
	uc := NewUsecase(repositoryMock)

	// When
	err := uc.Store(url)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, errMock, err)
}
