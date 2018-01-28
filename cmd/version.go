package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/morikuni/aec"
)

func newVersionCmd(args []string) *cobra.Command {

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "get version",
	}

	return versionCmd
}

func printLogo() {
	figletColoured := aec.BlueF.Apply(taasLogo)
	if runtime.GOOS == "windows" {
		figletColoured = aec.GreenF.Apply(taasLogo)
	}
	fmt.Printf(figletColoured)
}
