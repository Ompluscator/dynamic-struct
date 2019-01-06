package dynamicstruct

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/go-playground/form"
	"github.com/leebenson/conform"
	"gopkg.in/go-playground/validator.v9"
)

func TestNewBuilder(t *testing.T) {
	value := NewStruct()

	builder, ok := value.(*builderImpl)
	if !ok {
		t.Errorf(`TestNewBuilder - expected instance of *builder got %#v`, value)
	}

	if builder.fields == nil {
		t.Error(`TestNewBuilder - expected instance of *map[string]*fieldConfig got nil`)
	}

	if len(builder.fields) > 0 {
		t.Errorf(`TestNewBuilder - expected length of fields map to be 0 got %d`, len(builder.fields))
	}
}

func TestExtendStruct(t *testing.T) {
	value := ExtendStruct(struct {
		Field int `key:"value"`
	}{})

	builder, ok := value.(*builderImpl)
	if !ok {
		t.Errorf(`TestExtendStruct - expected instance of *builder got %#v`, value)
	}

	if builder.fields == nil {
		t.Error(`TestExtendStruct - expected instance of *map[string]*fieldConfig got nil`)
	}

	if len(builder.fields) != 1 {
		t.Errorf(`TestExtendStruct - expected length of fields map to be 1 got %d`, len(builder.fields))
	}

	field, ok := builder.fields["Field"]
	if !ok {
		t.Error(`TestExtendStruct - expected to have field "Field"`)
	}

	expected := &fieldConfigImpl{
		typ: 0,
		tag: `key:"value"`,
	}

	if !reflect.DeepEqual(field, expected) {
		t.Errorf(`TestExtendStruct - expected field to be %#v got %#v`, expected, field)
	}
}

func TestMergeStructs(t *testing.T) {
	value := MergeStructs(
		struct {
			FieldOne int `keyOne:"valueOne"`
		}{},
		struct {
			FieldTwo string `keyTwo:"valueTwo"`
		}{},
	)

	builder, ok := value.(*builderImpl)
	if !ok {
		t.Errorf(`TestMergeStructs - expected instance of *builder got %#v`, value)
	}

	if builder.fields == nil {
		t.Error(`TestMergeStructs - expected instance of *map[string]*fieldConfig got nil`)
	}

	if len(builder.fields) != 2 {
		t.Errorf(`TestMergeStructs - expected length of fields map to be 1 got %d`, len(builder.fields))
	}

	fieldOne, ok := builder.fields["FieldOne"]
	if !ok {
		t.Error(`TestMergeStructs - expected to have field "FieldOne"`)
	}

	expectedOne := &fieldConfigImpl{
		typ: 0,
		tag: `keyOne:"valueOne"`,
	}

	if !reflect.DeepEqual(fieldOne, expectedOne) {
		t.Errorf(`TestMergeStructs - expected field "FieldOne" to be %#v got %#v`, expectedOne, fieldOne)
	}

	fieldTwo, ok := builder.fields["FieldTwo"]
	if !ok {
		t.Error(`TestMergeStructs - expected to have field "FieldTwo"`)
	}

	expectedTwo := &fieldConfigImpl{
		typ: "",
		tag: `keyTwo:"valueTwo"`,
	}

	if !reflect.DeepEqual(fieldTwo, expectedTwo) {
		t.Errorf(`TestMergeStructs - expected field "FieldTwo" to be %#v got %#v`, expectedTwo, fieldTwo)
	}
}

func TestBuilderImpl_AddField(t *testing.T) {
	builder := &builderImpl{
		fields: map[string]*fieldConfigImpl{},
	}

	builder.AddField("Field", 1, `key:"value"`)

	field, ok := builder.fields["Field"]
	if !ok {
		t.Error(`TestBuilder_AddField - expected to have field "Field"`)
	}

	expected := &fieldConfigImpl{
		typ: 1,
		tag: `key:"value"`,
	}

	if !reflect.DeepEqual(field, expected) {
		t.Errorf(`TestExtendStruct - expected field to be %#v got %#v`, expected, field)
	}
}

func TestBuilderImpl_RemoveField(t *testing.T) {
	builder := &builderImpl{
		fields: map[string]*fieldConfigImpl{
			"Field": {
				tag: `key:"value"`,
				typ: 1,
			},
		},
	}

	builder.RemoveField("Field")

	if _, ok := builder.fields["Field"]; ok {
		t.Error(`TestBuilder_RemoveField - expected not to have field "Field"`)
	}
}

func TestBuilderImpl_HasField(t *testing.T) {
	builder := &builderImpl{
		fields: map[string]*fieldConfigImpl{},
	}

	if builder.HasField("Field") {
		t.Error(`TestBuilder_HasField - expected not to have field "Field"`)
	}

	builder = &builderImpl{
		fields: map[string]*fieldConfigImpl{
			"Field": {
				tag: `key:"value"`,
				typ: 1,
			},
		},
	}

	if !builder.HasField("Field") {
		t.Error(`TestBuilder_HasField - expected to have field "Field"`)
	}
}

func TestBuilderImpl_GetField(t *testing.T) {
	builder := &builderImpl{
		fields: map[string]*fieldConfigImpl{
			"Field": {
				tag: `key:"value"`,
				typ: 1,
			},
		},
	}

	value := builder.GetField("Field")

	field, ok := value.(*fieldConfigImpl)
	if !ok {
		t.Errorf(`TestBuilder_GetField - expected instance of *fieldConfig got %#v`, value)
	}

	expected := &fieldConfigImpl{
		typ: 1,
		tag: `key:"value"`,
	}

	if !reflect.DeepEqual(field, expected) {
		t.Errorf(`TestExtendStruct - expected field to be %#v got %#v`, expected, field)
	}
}

func TestFieldConfigImpl_SetTag(t *testing.T) {
	field := &fieldConfigImpl{}

	field.SetTag(`key:"value"`)

	if field.tag != `key:"value"` {
		t.Errorf(`TestFieldConfigImpl_SetTag - expected tag to be "%s" got "%s"`, `key:"value"`, field.tag)
	}
}

func TestFieldConfigImpl_SetType(t *testing.T) {
	field := &fieldConfigImpl{}

	field.SetType(1000)

	if field.typ != 1000 {
		t.Errorf(`TestFieldConfigImpl_SetType - expected type to be as for %#v got %#v`, 1000, field.typ)
	}
}

func TestBuilderImpl_Build_JSON(t *testing.T) {
	builder := getBuilder()

	builder.GetField("Integer").SetTag(`json:"int"`)
	builder.GetField("Text").SetTag(`json:"someText"`)
	builder.GetField("Float").SetTag(`json:"double"`)
	builder.GetField("Anonymous").SetTag(`json:"-"`)
	value := builder.Build()

	data := []byte(`
{
	"int": 123,
	"someText": "example",
	"double": 123.45,
	"Boolean": true,
	"Time": "2018-09-22T19:42:31+07:00",
	"Slice": [1, 2, 3],
	"Anonymous": "avoid to read",
	"NilFloat": 567
}
`)

	err := json.Unmarshal(data, &value)
	if err != nil {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected not to have error got %#v`, err)
	}

	valueOf := reflect.Indirect(reflect.ValueOf(value))

	if value := valueOf.FieldByName("Integer").Interface(); value != 123 {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "Integer" to be %#v got %#v`, 123, value)
	}

	if value := valueOf.FieldByName("Text").Interface(); value != "example" {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "Text" to be %#v got %#v`, "example", value)
	}

	if value := valueOf.FieldByName("Float").Interface(); value != 123.45 {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "Float" to be %#v got %#v`, 123.45, value)
	}

	if value := valueOf.FieldByName("Boolean").Interface(); value != true {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "Boolean" to be %#v got %#v`, true, value)
	}

	expected, err := time.Parse(time.RFC3339, "2018-09-22T19:42:31+07:00")
	if err != nil {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected not to have error got %#v`, err)
	}

	if value, ok := valueOf.FieldByName("Boolean").Interface().(time.Time); ok && value != expected {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "Time" to be %#v got %#v`, expected, value)
	}

	if value := valueOf.FieldByName("Slice").Interface(); !reflect.DeepEqual(value, []int{1, 2, 3}) {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "Slice" to be %#v got %#v`, []int{1, 2, 3}, value)
	}

	if value := valueOf.FieldByName("NilInteger").Interface(); value != (*int)(nil) {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "NilInteger" to be nil got %#v`, value)
	}

	if value := valueOf.FieldByName("NilText").Interface(); value != (*string)(nil) {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "NilText" to be nil got %#v`, value)
	}

	if value := valueOf.FieldByName("NilBoolean").Interface(); value != (*bool)(nil) {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "NilBoolean" to be nil got %#v`, value)
	}

	if value, ok := valueOf.FieldByName("NilFloat").Interface().(*float64); !ok || *value != 567.0 {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "NilFloat" to be %#v got %#v`, 567, value)
	}

	if value := valueOf.FieldByName("NilTime").Interface(); value != (*time.Time)(nil) {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "NilTime" to be nil got %#v`, value)
	}

	if value := valueOf.FieldByName("Anonymous").Interface(); value != "" {
		t.Errorf(`TestBuilderImpl_Build_JSON - expected field "Anonymous" to be %#v got %#v`, "", value)
	}
}

func TestBuilderImpl_Build_Form(t *testing.T) {
	builder := getBuilder()

	builder.GetField("Integer").SetTag(`form:"int"`)
	builder.GetField("Text").SetTag(`form:"someText" conform:"trim"`)
	builder.GetField("Float").SetTag(`form:"double"`)
	builder.GetField("Anonymous").SetTag(`form:"-"`)
	value := builder.Build()

	data := url.Values{
		"int":       []string{"123"},
		"someText":  []string{" example "},
		"double":    []string{"123.45"},
		"Boolean":   []string{"on"},
		"Time":      []string{"2018-09-22T19:42:31+07:00"},
		"Slice":     []string{"1", "2", "3"},
		"Anonymous": []string{"avoid to read"},
		"NilFloat":  []string{"567"},
	}

	decoder := form.NewDecoder()
	err := decoder.Decode(&value, data)
	if err != nil {
		t.Errorf(`TestBuilderImpl_Build_Form - expected not to have error got %#v`, err)
	}

	err = conform.Strings(value)
	if err != nil {
		t.Errorf(`TestBuilderImpl_Build_Form - expected not to have error got %#v`, err)
	}

	valueOf := reflect.Indirect(reflect.ValueOf(value))

	if value := valueOf.FieldByName("Integer").Interface(); value != 123 {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "Integer" to be %#v got %#v`, 123, value)
	}

	if value := valueOf.FieldByName("Text").Interface(); value != "example" {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "Text" to be %#v got %#v`, "example", value)
	}

	if value := valueOf.FieldByName("Float").Interface(); value != 123.45 {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "Float" to be %#v got %#v`, 123.45, value)
	}

	if value := valueOf.FieldByName("Boolean").Interface(); value != true {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "Boolean" to be %#v got %#v`, true, value)
	}

	expected, err := time.Parse(time.RFC3339, "2018-09-22T19:42:31+07:00")
	if err != nil {
		t.Errorf(`TestBuilderImpl_Build_Form - expected not to have error got %#v`, err)
	}

	if value, ok := valueOf.FieldByName("Boolean").Interface().(time.Time); ok && value != expected {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "Time" to be %#v got %#v`, expected, value)
	}

	if value := valueOf.FieldByName("Slice").Interface(); !reflect.DeepEqual(value, []int{1, 2, 3}) {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "Slice" to be %#v got %#v`, []int{1, 2, 3}, value)
	}

	if value := valueOf.FieldByName("NilInteger").Interface(); value != (*int)(nil) {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "NilInteger" to be nil got %#v`, value)
	}

	if value := valueOf.FieldByName("NilText").Interface(); value != (*string)(nil) {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "NilText" to be nil got %#v`, value)
	}

	if value := valueOf.FieldByName("NilBoolean").Interface(); value != (*bool)(nil) {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "NilBoolean" to be nil got %#v`, value)
	}

	if value, ok := valueOf.FieldByName("NilFloat").Interface().(*float64); !ok || *value != 567.0 {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "NilFloat" to be %#v got %#v`, 567, value)
	}

	if value := valueOf.FieldByName("NilTime").Interface(); value != (*time.Time)(nil) {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "NilTime" to be nil got %#v`, value)
	}

	if value := valueOf.FieldByName("Anonymous").Interface(); value != "" {
		t.Errorf(`TestBuilderImpl_Build_Form - expected field "Anonymous" to be %#v got %#v`, "", value)
	}
}

func TestBuilderImpl_Build_Validate(t *testing.T) {
	builder := getBuilder()

	builder.GetField("Integer").SetTag(`validate:"gt=0"`)
	builder.GetField("Float").SetTag(`validate:"gte=0"`)
	builder.GetField("Text").SetTag(`validate:"required"`)
	value := builder.Build()

	err := validator.New().Struct(value)
	if err == nil {
		t.Errorf(`TestBuilderImpl_Build_Validate - expected to have error got %#v`, err)
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		t.Errorf(`TestBuilderImpl_Build_Validate - expected instance of *validator.ValidationErrors got %#v`, value)
	}

	for _, fieldError := range validationErrors {
		fieldError.Tag()
	}

	fieldErrors := []validator.FieldError(validationErrors)
	if len(fieldErrors) != 2 {
		t.Errorf(`TestBuilderImpl_Build_Validate - expected field errors to have length %#v got %#v`, 2, len(fieldErrors))
	}

	checkedInteger := false
	checkedText := false
	for _, fieldError := range fieldErrors {
		if fieldError.Field() == "Integer" {
			checkedInteger = true

			if fieldError.Tag() != "gt" {
				t.Errorf(`TestBuilderImpl_Build_Validate - expected tag of field error to be %#v got %#v`, "gt", fieldError.Tag())
			}

			if fieldError.Param() != "0" {
				t.Errorf(`TestBuilderImpl_Build_Validate - expected param of field error to be %#v got %#v`, "0", fieldError.Param())
			}
		} else if fieldError.Field() == "Text" {
			checkedText = true

			if fieldError.Tag() != "required" {
				t.Errorf(`TestBuilderImpl_Build_Validate - expected tag of field error to be %#v got %#v`, "required", fieldError.Tag())
			}
		}
	}

	if !checkedInteger {
		t.Error(`TestBuilderImpl_Build_Validate - expected to have field errors for field "Integer"`)
	}

	if !checkedText {
		t.Error(`TestBuilderImpl_Build_Validate - expected to have field errors for field "Text"`)
	}
}

func getBuilder() Builder {
	integer := 0
	str := ""
	float := 0.0
	boolean := false

	return NewStruct().
		AddField("Integer", integer, "").
		AddField("Text", str, "").
		AddField("Float", float, "").
		AddField("Boolean", boolean, "").
		AddField("Time", time.Time{}, "").
		AddField("Slice", []int{}, "").
		AddField("NilInteger", &integer, "").
		AddField("NilText", &str, "").
		AddField("NilFloat", &float, "").
		AddField("NilBoolean", &boolean, "").
		AddField("NilTime", &time.Time{}, "").
		AddField("Anonymous", "", "")
}
