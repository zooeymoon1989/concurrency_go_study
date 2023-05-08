package goconvey

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T)  {
	Convey("Given some integer with a staring value", t , func() {
		x := 1
		Convey("When the integer is incremented", func() {
			x++
			Convey("The value should be greater by one" , func() {
				So(x , ShouldEqual , 2)
			})
		})
	})
}
