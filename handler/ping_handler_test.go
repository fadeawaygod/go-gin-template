package handler

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestGetPingShouldReturnPong(t *testing.T) {
	c := mock.Mock
	GetPing(c)
}
