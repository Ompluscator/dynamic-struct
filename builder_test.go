package dynamicstruct

import (
	"reflect"
	"testing"
)

func TestNewStruct(t *testing.T) {
	value := NewStruct()

	builder, ok := value.(*builderImpl)
	if !ok {
		t.Errorf(`TestNewStruct - expected instance of *builder got %#v`, value)
	}

	if builder.fields == nil {
		t.Error(`TestNewStruct - expected instance of *map[string]*fieldConfig got nil`)
	}

	if len(builder.fields) > 0 {
		t.Errorf(`TestNewStruct - expected length of fields map to be 0 got %d`, len(builder.fields))
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

func TestExtendStructWithSettableFields(t *testing.T) {
	value, err := ExtendStructWithSettableFields(&struct {
		f     int
		Field int `key:"value"`
	}{})

	if err != nil {
		t.Errorf(`TestExtendStructWithSettableFields - expected nil error got %#v`, err)
	}

	builder, ok := value.(*builderImpl)
	if !ok {
		t.Errorf(`TestExtendStructWithSettableFields - expected instance of *builder got %#v`, value)
	}

	if builder.fields == nil {
		t.Error(`TestExtendStructWithSettableFields - expected instance of *map[string]*fieldConfig got nil`)
	}

	if len(builder.fields) != 1 {
		t.Errorf(`TestExtendStructWithSettableFields - expected length of fields map to be 1 got %d`, len(builder.fields))
	}

	field, ok := builder.fields["Field"]
	if !ok {
		t.Error(`TestExtendStructWithSettableFields - expected to have field "Field"`)
	}

	expected := &fieldConfigImpl{
		typ: 0,
		tag: `key:"value"`,
	}

	if !reflect.DeepEqual(field, expected) {
		t.Errorf(`TestExtendStructWithExportedFields - expected field to be %#v got %#v`, expected, field)
	}
}

func TestMergeStructsWithSettableFields(t *testing.T) {
	value, err := MergeStructsWithSettableFields(
		&struct {
			FieldOne int `keyOne:"valueOne"`
		}{},
		&struct {
			FieldTwo string `keyTwo:"valueTwo"`
		}{},
	)

	if err != nil {
		t.Errorf(`TestMergeStructsWithSettableFields - expected nil error got %#v`, err)
	}

	builder, ok := value.(*builderImpl)
	if !ok {
		t.Errorf(`TestMergeStructsWithSettableFields - expected instance of *builder got %#v`, value)
	}

	if builder.fields == nil {
		t.Error(`TestMergeStructsWithSettableFields - expected instance of *map[string]*fieldConfig got nil`)
	}

	if len(builder.fields) != 2 {
		t.Errorf(`TestMergeStructsWithSettableFields - expected length of fields map to be 1 got %d`, len(builder.fields))
	}

	fieldOne, ok := builder.fields["FieldOne"]
	if !ok {
		t.Error(`TestMergeStructsWithSettableFields - expected to have field "FieldOne"`)
	}

	expectedOne := &fieldConfigImpl{
		typ: 0,
		tag: `keyOne:"valueOne"`,
	}

	if !reflect.DeepEqual(fieldOne, expectedOne) {
		t.Errorf(`TestMergeStructsWithSettableFields - expected field "FieldOne" to be %#v got %#v`, expectedOne, fieldOne)
	}

	fieldTwo, ok := builder.fields["FieldTwo"]
	if !ok {
		t.Error(`TestMergeStructsWithSettableFields - expected to have field "FieldTwo"`)
	}

	expectedTwo := &fieldConfigImpl{
		typ: "",
		tag: `keyTwo:"valueTwo"`,
	}

	if !reflect.DeepEqual(fieldTwo, expectedTwo) {
		t.Errorf(`TestMergeStructsWithSettableFields - expected field "FieldTwo" to be %#v got %#v`, expectedTwo, fieldTwo)
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

	undefined := builder.GetField("Undefined")
	if undefined != nil {
		t.Errorf(`TestBuilder_GetField - expected nil got %#v`, value)
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
