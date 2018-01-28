package version

import (
	"fmt"
	"testing"

	"github.com/zanetworker/taas/pkg/testutils"
)

func TestVersion(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		expected string
	}{
		{
			name:     "Empty Version Means BuildVersion Returns Dev",
			id:       "EmptyVersionMeansBuildVersionReturnsDev",
			expected: "dev",
		},
		{
			name:     "Version Returned FromFrom Build Version",
			id:       "VersionReturnedFromBuildVersion",
			expected: "testing-manual",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output, errMsg string

			switch tt.id {
			case "EmptyVersionMeansBuildVersionReturnsDev":
				output = BuildVersion()
				errMsg = fmt.Sprintf("Version is not from Build - want: %s, got: %s", tt.expected, output)

			case "VersionReturnedFromBuildVersion":
				Version = "testing-manual"
				output = BuildVersion()
				errMsg = fmt.Sprintf("Version is not from Build - want: %s, got: %s", tt.expected, output)
			}

			testutils.Equals(t, tt.expected, output, errMsg)
		})
	}
}
