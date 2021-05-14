package parser

import (
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListParser(t *testing.T) {
	Convey("Scenario: Test word list parsing", t, func() {
		file, _ := ioutil.ReadFile("letra_a.html")
		content := string(file)
		logrus.Infof("Words list mock file contents: %s", content)

		Convey("Given word list parser instance", func() {
			instance := NewListParser()

			Convey("When parser method is invoked", func() {
				result := instance.Parse(content)

				Convey("Then the list of words should be parsed", func() {
					So(len(result), ShouldBeGreaterThan, 20)
					So(result[0].Name, ShouldEqual, "aarónico")
					So(result[0].URL, ShouldEqual, "https://dle.rae.es/aarónico")
					So(result[1].Name, ShouldEqual, "aaronita")
					So(result[1].URL, ShouldEqual, "https://dle.rae.es/aaronita")
					So(result[2].Name, ShouldEqual, "ababa")
					So(result[2].URL, ShouldEqual, "https://dle.rae.es/ababa")

					Convey("And it should not contain conjugation words", func() {
						for _, r := range result {
							So(r.Name, ShouldNotEqual, "ababilla")
						}
					})
				})
			})
		})
	})
}
