package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {

	_, err := Init()
	assert.NoError(t, err)
}
