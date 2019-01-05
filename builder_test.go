package dynamic_struct

import (
	"encoding/json"
	"github.com/go-playground/form"
	"github.com/leebenson/conform"
	"gopkg.in/go-playground/validator.v9"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

type (
	BuilderTestSuite struct {
		suite.Suite

		builder Builder
	}
)

func TestBuilderTestSuite(t *testing.T) {
	suite.Run(t, &BuilderTestSuite{})
}

func (t *BuilderTestSuite) SetupTest() {
	integer := 0
	str := ""
	float := 0.0
	boolean := false

	t.builder = NewBuilder().
		AddField("Integer", integer, "").
		AddField("Text", str, "").
		AddField("Float", float, "").
		AddField("Boolean", boolean, "").
		AddField("NilInteger", &integer, "").
		AddField("NilText", &str, "").
		AddField("NilFloat", &float, "").
		AddField("NilBoolean", &boolean, "").
		AddField("Slice", []int{}, "").
		AddField("Anonymous", "", "")
}

func (t *BuilderTestSuite) TestJson() {
	t.builder.GetField("Integer").SetTag(`json:"int"`)
	t.builder.GetField("Text").SetTag(`json:"someText"`)
	t.builder.GetField("Float").SetTag(`json:"double"`)
	t.builder.GetField("Anonymous").SetTag(`json:"-"`)
	value := t.builder.Build()

	data := []byte(`
{
	"int": 123,
	"someText": "example",
	"double": 123.45,
	"Boolean": true,
	"Slice": [1, 2, 3],
	"Anonymous": "avoid to read",
	"NilFloat": 567
}
`)

	err := json.Unmarshal(data, &value)
	t.NoError(err)

	var result map[string]interface{}
	err = NewReader(value).MapTo(&result)
	t.NoError(err)

	float := 567.0
	t.Equal(map[string]interface{}{
		"int": 123.0,
		"someText": "example",
		"double": 123.45,
		"Boolean": true,
		"Slice": []interface{}{1.0, 2.0, 3.0},
		"NilInteger": nil,
		"NilFloat": float,
		"NilBoolean": nil,
		"NilText": nil,
	}, result)

}

func (t *BuilderTestSuite) TestFormAndConform() {
	t.builder.GetField("Integer").SetTag(`form:"int"`)
	t.builder.GetField("Text").SetTag(`form:"someText" conform:"trim"`)
	t.builder.GetField("Float").SetTag(`form:"double"`)
	t.builder.GetField("Anonymous").SetTag(`form:"-"`)
	value := t.builder.Build()

	data := url.Values{
		"int":       []string{"123"},
		"someText":  []string{" example "},
		"double":    []string{"123.45"},
		"Boolean":   []string{"on"},
		"Slice":     []string{"1", "2", "3"},
		"Anonymous": []string{"avoid to read"},
		"NilFloat":  []string{"567"},
	}

	decoder := form.NewDecoder()
	err := decoder.Decode(&value, data)
	t.NoError(err)

	err = conform.Strings(value)
	t.NoError(err)

	var result map[string]interface{}
	err = NewReader(value).MapTo(&result)
	t.NoError(err)

	float := 567.0
	t.Equal(map[string]interface{}{
		"Integer": 123.0,
		"Text": "example",
		"Float": 123.45,
		"Boolean": true,
		"Anonymous": "",
		"Slice": []interface{}{1.0, 2.0, 3.0},
		"NilInteger": nil,
		"NilFloat": float,
		"NilBoolean": nil,
		"NilText": nil,
	}, result)
}

func (t *BuilderTestSuite) TestValidate() {
	t.builder.GetField("Integer").SetTag(`validate:"gt=0"`)
	t.builder.GetField("Text").SetTag(`validate:"required"`)
	value := t.builder.Build()

	err := validator.New().Struct(value)
	t.Error(err)

	validationErrors, ok := err.(validator.ValidationErrors)
	t.True(ok)
	t.Len(validationErrors, 2)
}
