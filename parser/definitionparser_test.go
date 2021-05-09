package parser

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefinitionParser(t *testing.T) {
	Convey("Scenario: test word definitions parsing", t, func() {
		file, _ := ioutil.ReadFile("ababillarse.html")
		content := string(file)

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

	Convey("Scenario: test word with no or wrong definition", t, func() {
		file, _ := ioutil.ReadFile("abachar.html")
		content := string(file)

		Convey("Given a file with no definition and a parser instance", func() {
			parser := NewDefinitionParser()

			Convey("When parse method is invoked with file contents", func() {
				result := parser.Parse(content)

				Convey("Then the extracted definition should be empty", func() {
					So(result, ShouldEqual, "")
				})
			})
		})
	})

	Convey("Scenario: test word with broken definition", t, func() {
		file, _ := ioutil.ReadFile("abalada.html")
		content := string(file)

		Convey("Given a definition parser instance", func() {
			parser := NewDefinitionParser()

			Convey("When parse method is invoked with file having corrupt definition", func() {
				result := parser.Parse(content)

				Convey("Then the definition should be empty", func() {
					So(result, ShouldEqual, "")
				})
			})
		})
	})

	Convey("Scenerio: test word with short definition", t, func() {
		file, _ := ioutil.ReadFile("aba.html")
		content := string(file)

		Convey("Given a definition parser instance", func() {
			parser := NewDefinitionParser()

			Convey("When parse method is invoked with file having short definition", func() {
				result := parser.Parse(content)

				Convey("Then the definition should be empty", func() {
					So(result, ShouldEqual, "")
				})
			})
		})
	})

	Convey("Scenario: test word definition with parentheses", t, func() {
		file, _ := ioutil.ReadFile("abadiado.html")
		content := string(file)

		Convey("Given a definition parser instance", func() {
			parser := NewDefinitionParser()

			Convey("When parse method is invoked with file having definition with parentheses", func() {
				result := parser.Parse(content)

				Convey("Then definition should not have the part between parentheses", func() {
					So(result, ShouldEqual, "")
				})
			})
		})
	})

	Convey("Scenario: test missing word definition", t, func() {
		file, _ := ioutil.ReadFile("abadengo.html")
		content := string(file)

		Convey("Given a definition parser instance", func() {
			parser := NewDefinitionParser()

			Convey("When parse method is invoked with file having definition with parentheses", func() {
				result := parser.Parse(content)

				Convey("Then definition should not have the part between parentheses", func() {
					So(result, ShouldEqual, "")
				})
			})
		})
	})
}
