package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/zanetworker/opentaas/pkg/globalutils"
	"github.com/zanetworker/opentaas/pkg/testutils"
)

//TODO: implement tests for goss command

func TestGetGossConfigDir(t *testing.T) {
	dirPathTrailExpected := "/opentaas/configs/goss"
	getGossDirPath := globalutils.GetDir("config_goss")

	testutils.Assert(t, strings.Contains(getGossDirPath, dirPathTrailExpected),
		fmt.Sprintf("goss config dir is not being fetched correctly (%s)", getGossDirPath))
}
