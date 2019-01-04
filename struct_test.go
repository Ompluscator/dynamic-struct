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
	StructTestSuite struct {
		suite.Suite

		builder Builder
	}
)

func TestStructTestSuite(t *testing.T) {
	suite.Run(t, &StructTestSuite{})
}

func (t *StructTestSuite) SetupTest() {
	t.builder = NewBuilder().
		AddField("Integer", 0, "").
		AddField("Text", "", "").
		AddField("Float", 0.0, "").
		AddField("Boolean", false, "").
		AddField("Slice", []int{}, "").
		AddField("Anonymous", "", "")
}

func (t *StructTestSuite) TestJson() {
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
	"Anonymous": "avoid to read"
}
`)

	err := json.Unmarshal(data, &value)
	t.NoError(err)

	vReader := NewReader(value)
	t.Equal(123, vReader.Int("Integer"))
	t.Equal("example", vReader.String("Text"))
	t.Equal(123.45, vReader.Float64("Float"))
	t.Equal(true, vReader.Bool("Boolean"))
	t.Equal("", vReader.String("Anonymous"))
}

func (t *StructTestSuite) TestFormAndConform() {
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
	}

	decoder := form.NewDecoder()
	err := decoder.Decode(&value, data)
	t.NoError(err)

	err = conform.Strings(value)
	t.NoError(err)

	vReader := NewReader(value)
	t.Equal(123, vReader.Int("Integer"))
	t.Equal("example", vReader.String("Text"))
	t.Equal(123.45, vReader.Float64("Float"))
	t.Equal(true, vReader.Bool("Boolean"))
	t.Equal("", vReader.String("Anonymous"))
}

func (t *StructTestSuite) TestValidate() {
	t.builder.GetField("Integer").SetTag(`validate:"gt=0"`)
	t.builder.GetField("Text").SetTag(`validate:"required"`)
	value := t.builder.Build()

	err := validator.New().Struct(value)
	t.Error(err)

	validationErrors, ok := err.(validator.ValidationErrors)
	t.True(ok)
	t.Len(validationErrors, 2)
}