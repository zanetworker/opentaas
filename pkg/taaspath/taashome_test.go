package taaspath_test

import (
	"testing"

	"github.com/zanetworker/taas/pkg/taaspath"
	"github.com/zanetworker/taas/pkg/testutils"
)

func TestTaasHome(t *testing.T) {
	hh := taaspath.Home("/r")
	testutils.Equals(t, hh.String(), "/r")
	testutils.Equals(t, hh.TLSCaCert(), "/r/ca.pem")
	testutils.Equals(t, hh.TLSCert(), "/r/cert.pem")
	testutils.Equals(t, hh.TLSKey(), "/r/key.pem")
}

func TestTaasHomeExpand(t *testing.T) {
	conditionToAssert := taaspath.Home("$HOME").String() != "$HOME"
	testutils.Assert(t, conditionToAssert, "home variables is not expanded correctly")
}
