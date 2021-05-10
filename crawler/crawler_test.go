package crawler

import (
	"dictio-scrapper/model"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

type mockTTP struct {
	mock.Mock
	HttpGetter
}

type mockListParser struct {
	mock.Mock
}

type mockDefinitionParser struct {
	mock.Mock
}

type mockDb struct {
	mock.Mock
}

func (m *mockTTP) Get(url string) (string, error) {
	args := m.Called(url)
	return args.String(0), args.Error(1)
}

func (m *mockListParser) Parse(content string) []model.Word {
	args := m.Called(content)
	return args.Get(0).([]model.Word)
}

func (m *mockDefinitionParser) Parse(content string) string {
	args := m.Called(content)
	return args.String(0)
}

func (m *mockDb) Save(data model.Entry) error {
	args := m.Called(data)
	return args.Error(0)
}

func TestCrawlingProcess(t *testing.T) {
	Convey("Scenario: test crawling process", t, func() {
		httpGetter := new(mockTTP)
		listParser := new(mockListParser)
		definitionParser := new(mockDefinitionParser)
		db := new(mockDb)

		httpGetter.On("Get", "list-A.com").Return("<html>words list</html>", nil)
		httpGetter.On("Get", "definition.com").Return("<html>word definition</html>", nil)

		listParser.On("Parse", "<html>words list</html>").Return([]model.Word{
			model.Word{Name: "foo", URL: "definition.com"},
		})

		definitionParser.On("Parse", "<html>word definition</html>").Return("lorem ipsum dolor sit amet")

		db.On("Save", mock.AnythingOfType("model.Entry")).Return(nil)

		Convey("Given a crawler instance", func() {
			instance := New(httpGetter, listParser, definitionParser, db, []string{"A"})

			Convey("When crawling process is invoked", func() {
				err := instance.Process("list-${LETTER}.com")

				Convey("Then crawling is executed", func() {
					So(err, ShouldBeNil)
					httpGetter.AssertNumberOfCalls(t, "Get", 2)
					httpGetter.AssertCalled(t, "Get", "list-A.com")
					httpGetter.AssertCalled(t, "Get", "definition.com")

					listParser.AssertNumberOfCalls(t, "Parse", 1)
					listParser.AssertCalled(t, "Parse", "<html>words list</html>")

					definitionParser.AssertNumberOfCalls(t, "Parse", 1)
					definitionParser.AssertCalled(t, "Parse", "<html>word definition</html>")

					db.AssertNumberOfCalls(t, "Save", 1)
					db.AssertCalled(t, "Save", model.Entry{
						Word:       "foo",
						Definition: "lorem ipsum dolor sit amet",
						Letters:    []string{"f"},
					})
				})
			})
		})
	})

	Convey("Scenario: test crawling process with special initial letter", t, func() {
		httpGetter := new(mockTTP)
		listParser := new(mockListParser)
		definitionParser := new(mockDefinitionParser)
		db := new(mockDb)

		httpGetter.On("Get", "list-A.com").Return("<html>words list</html>", nil)
		httpGetter.On("Get", "definition.com").Return("<html>word definition</html>", nil)

		listParser.On("Parse", "<html>words list</html>").Return([]model.Word{
			model.Word{Name: "ámbar", URL: "definition.com"},
		})

		definitionParser.On("Parse", "<html>word definition</html>").Return("lorem ipsum dolor sit amet")

		db.On("Save", mock.AnythingOfType("model.Entry")).Return(nil)

		Convey("Given a crawler instance", func() {
			instance := New(httpGetter, listParser, definitionParser, db, []string{"A"})

			Convey("When crawling process is invoked", func() {
				err := instance.Process("list-${LETTER}.com")

				Convey("Then crawling is executed", func() {
					So(err, ShouldBeNil)
					httpGetter.AssertNumberOfCalls(t, "Get", 2)
					httpGetter.AssertCalled(t, "Get", "list-A.com")
					httpGetter.AssertCalled(t, "Get", "definition.com")

					listParser.AssertNumberOfCalls(t, "Parse", 1)
					listParser.AssertCalled(t, "Parse", "<html>words list</html>")

					definitionParser.AssertNumberOfCalls(t, "Parse", 1)
					definitionParser.AssertCalled(t, "Parse", "<html>word definition</html>")

					db.AssertNumberOfCalls(t, "Save", 1)
					db.AssertCalled(t, "Save", model.Entry{
						Word:       "ámbar",
						Definition: "lorem ipsum dolor sit amet",
						Letters:    []string{"a"},
					})
				})
			})
		})
	})
}
