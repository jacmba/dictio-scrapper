package parser

import (
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDefinitionParser(t *testing.T) {
	Convey("Scenario: test word definitions parsing", t, func() {
		file, _ := ioutil.ReadFile("ababillarse.html")
		content := string(file)
		logrus.Infof("Mock file contents: %s", content)

		Convey("Given a definition parser instance", func() {
			parser := NewDefinitionParser()

			Convey("When parse method is invoked with file contents", func() {
				result := parser.Parse(content)

				Convey("Then the definition of the word should be extracted", func() {
					So(result, ShouldEqual, "Dicho de un animal: Enfermar de la babilla")
				})
			})
		})
	})
}
