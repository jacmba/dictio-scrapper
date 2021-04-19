package config

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfiguration(t *testing.T) {
	Convey("Scenario: load configuration", t, func() {
		Convey("When loading default configuration", func() {
			LoadConfig()
			Convey("Then default values from yaml file should be loaded", func() {
				So(GlobalConfig.URL, ShouldEqual, "https://www.listapalabras.com/palabras-con-${LETTER}-lista-completa.php")
			})
		})

		Convey("When loading configuratio with env overriden value", func() {
			os.Setenv("URL", "my_mock_url")
			LoadConfig()
			Convey("Then overriden value should be loaded", func() {
				So(GlobalConfig.URL, ShouldEqual, "my_mock_url")
			})
		})
	})
}
