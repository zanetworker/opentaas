package log

import (
	"testing"

	"github.com/zanetworker/opentaas/pkg/testutils"
)

func TestGetFileNameCapitalized(t *testing.T) {
	input := "test/parent/file"
	expected := "PARENT|FILE"
	actual := getFileNameCapitalized(input)

	testutils.Equals(t, expected, actual, "getFileNameCapitalized is not doing it's job")
}

// func TestLogWrapper(t *testing.T) {
// 	tests := []struct {
// 		caseName     string
// 		messageToLog string
// 		expected     string
// 	}{
// 		{
// 			caseName:     "Test Error Log",
// 			messageToLog: "dummy message",
// 		},
// 	}

// }
