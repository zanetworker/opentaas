package taaspath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaasHome(t *testing.T) {
	hh := Home("/r")
	assert := assert.New(t)

	assert.Equal(hh.String(), "/r", "theys should be equal")
	assert.Equal(hh.TLSCaCert(), "/r/ca.pem")
	assert.Equal(hh.TLSCert(), "/r/cert.pem")
	assert.Equal(hh.TLSKey(), "/r/key.pem")
}

func TestTaasHomeExpand(t *testing.T) {
	assert := assert.New(t)
	assert.NotEqual(Home("$HOME").String(), "$HOME")
}
