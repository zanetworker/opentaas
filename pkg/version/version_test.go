package version

import (
	"path/filepath"
	"testing"

	"github.com/zanetworker/taas/pkg/testutils"
)

func TestGetVersionHelper(t *testing.T) {
	versionFile := filepath.Base(getVersion())
	testutils.Assert(t, versionFile == "VERSION", "getVersion failed to return the path of the VERSION file")
}

// func TestVersion(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		id       string
// 		expected string
// 	}{
// 		{
// 			name:     "Empty Version Means BuildVersion Returns Dev",
// 			id:       "EmptyVersionMeansBuildVersionReturnsDev",
// 			expected: "dev",
// 		},
// 		{
// 			name:     "Version Returned FromFrom Build Version",
// 			id:       "VersionReturnedFromBuildVersion",
// 			expected: "testing-manual",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var actual, errMsg string

// 			switch tt.id {
// 			case "EmptyVersionMeansBuildVersionReturnsDev":
// 				actual = BuildVersion()
// 				errMsg = fmt.Sprintf("Version is not from Build - want: %s, got: %s", tt.expected, actual)

// 			case "VersionReturnedFromBuildVersion":
// 				Version = "testing-manual"
// 				actual = BuildVersion()
// 				errMsg = fmt.Sprintf("Version is not from Build - want: %s, got: %s", tt.expected, actual)
// 			}

// 			testutils.Equals(t, tt.expected, actual, errMsg)
// 		})
// 	}
// }
