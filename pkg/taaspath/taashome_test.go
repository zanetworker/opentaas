package taaspath_test

import (
	"testing"

	"github.com/zanetworker/opentaas/pkg/taaspath"
	"github.com/zanetworker/opentaas/pkg/testutils"
)

func TestTaasHome(t *testing.T) {
	hh := taaspath.Home("/r")
	testutils.Equals(t, hh.String(), "/r", "")
	testutils.Equals(t, hh.TLSCaCert(), "/r/ca.pem", "TLSCACert is not working")
	testutils.Equals(t, hh.TLSCert(), "/r/cert.pem", "TLSCert is not working")
	testutils.Equals(t, hh.TLSKey(), "/r/key.pem", "TLSCACert is not working")
}

func TestTaasHomeExpand(t *testing.T) {
	conditionToAssert := taaspath.Home("$HOME").String() != "$HOME"
	testutils.Assert(t, conditionToAssert, "home variables is not expanded correctly")
}
