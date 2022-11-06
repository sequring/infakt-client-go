package infakt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientModel(t *testing.T) {
	client := GetInfactClient()
	if assert.NotNil(t, client) {
		t.Log("Infakt client initialized")
	}
}
