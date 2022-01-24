package root

import (
	"bytes"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVersion(t *testing.T) {
	Convey("Calling the `ting version` command", t, func() {
		cmd := VersionCmd()
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.Execute()
		out, _ := ioutil.ReadAll(b)
		expected := "ting v0.1.0\n"

		Convey("The command should return correct version", func() {
			So(string(out), ShouldEqual, expected)
		})
	})
}
