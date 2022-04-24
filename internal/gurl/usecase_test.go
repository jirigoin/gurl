package gurl

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUseCase_GetAll_NotFound(t *testing.T) {
	// Given
	url := "some-url"
	file := "some file"
	controller := gomock.NewController(t)
	repositoryMock := NewMockRepositoryImpl(controller)
	repositoryMock.EXPECT().Store(url, file).Return()
	uc := NewUsecase(repositoryMock)

	// When
	apiError := uc.Store(url)

	// Then
	assert.Nil(t, apiError)
}
