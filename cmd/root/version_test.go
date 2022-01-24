package root

import (
	"testing"
)

func TestVersion(t *testing.T) {
	err := VersionCmd.Execute()

	if err != nil {
		t.Errorf("test failed: %s", err.Error())
	}
}
