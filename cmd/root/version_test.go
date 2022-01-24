package root

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVersion(t *testing.T) {
	Convey("Calling the `ting version` command", t, func() {

		err := VersionCmd().Execute()

		Convey("If no error", func() {
			So(err, ShouldEqual, nil)
		})

	})
}
