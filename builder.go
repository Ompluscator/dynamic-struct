package dynamicstruct

import (
	"errors"
	"reflect"
)

type (
	// Builder holds all fields' definitions for desired structs.
	// It gives opportunity to add or remove fields, or to edit
	// existing ones. In the end, it provides definition for
	// dynamic structs, used to create new instances.
	Builder interface {
		// AddField creates new struct's field.
		// It expects field's name, type and string.
		// Type is provided as an instance of some golang type.
		// Tag is provided as classical golang field tag.
		//
		// builder.AddField("SomeFloatField", 0.0, `json:"boolean" validate:"gte=10"`)
		//
		AddField(name string, typ interface{}, tag string) Builder
		// RemoveField removes existing struct's field.
		//
		// builder.RemoveField("SomeFloatField")
		//
		RemoveField(name string) Builder
		// HasField checks if struct has a field with a given name.
		//
		// if builder.HasField("SomeFloatField") { ...
		//
		HasField(name string) bool
		// GetField returns struct's field definition.
		// If there is no such field, it returns nil.
		// Usable to edit existing struct's field.
		//
		// field := builder.GetField("SomeFloatField")
		//
		GetField(name string) FieldConfig
		// Build returns definition for dynamic struct.
		// Definition can be used to create new instances.
		//
		// dStruct := builder.Build()
		//
		Build() DynamicStruct
	}

	// FieldConfig holds single field's definition.
	// It provides possibility to edit field's type and tag.
	FieldConfig interface {
		// SetType changes field's type.
		// Expected value is an instance of golang type.
		//
		// field.SetType([]int{})
		//
		SetType(typ interface{}) FieldConfig
		// SetTag changes fields's tag.
		// Expected value is an string which represents classical
		// golang tag.
		//
		// field.SetTag(`json:"slice"`)
		//
		SetTag(tag string) FieldConfig
	}

	// DynamicStruct contains defined dynamic struct.
	// This definition can't be changed anymore, once is built.
	// It provides a method for creating new instances of same defintion.
	DynamicStruct interface {
		// New provides new instance of defined dynamic struct.
		//
		// value := dStruct.New()
		//
		New() interface{}

		// NewSliceOfStructs provides new slice of defined dynamic struct, with 0 length and capacity.
		//
		// value := dStruct.NewSliceOfStructs()
		//
		NewSliceOfStructs() interface{}

		// New provides new map of defined dynamic struct with desired key type.
		//
		// value := dStruct.NewMapOfStructs("")
		//
		NewMapOfStructs(key interface{}) interface{}
	}

	builderImpl struct {
		fields map[string]*fieldConfigImpl
	}

	fieldConfigImpl struct {
		typ interface{}
		tag string
	}

	dynamicStructImpl struct {
		definition reflect.Type
	}
)

// NewStruct returns new clean instance of Builder interface
// for defining fresh dynamic struct.
//
// builder := dynamicstruct.NewStruct()
//
func NewStruct() Builder {
	return &builderImpl{
		fields: map[string]*fieldConfigImpl{},
	}
}

// ExtendStruct extends existing instance of struct and
// returns new instance of Builder interface.
//
// builder := dynamicstruct.MergeStructs(MyStruct{})
//
func ExtendStruct(value interface{}) Builder {
	return MergeStructs(value)
}

// MergeStructs merges a list of existing instances of structs and
// returns new instance of Builder interface.
//
// builder := dynamicstruct.MergeStructs(MyStructOne{}, MyStructTwo{}, MyStructThree{})
//
func MergeStructs(values ...interface{}) Builder {
	builder := NewStruct()

	for _, value := range values {
		valueOf := reflect.Indirect(reflect.ValueOf(value))
		typeOf := valueOf.Type()

		for i := 0; i < valueOf.NumField(); i++ {
			fval := valueOf.Field(i)
			ftyp := typeOf.Field(i)
			builder.AddField(ftyp.Name, fval.Interface(), string(ftyp.Tag))
		}
	}

	return builder
}

// ExtendStructWithSettableFields extends existing instance of struct and
// returns new instance of Builder interface.
//
// builder := dynamicstruct.ExtendStructWithSettableFields(MyStruct{})
//
// it will build only with exported and addressable fields,
// therefore it will not panic, even there are unexported or unaddressable fields in the struct,
// input value must be a pointer to struct, otherwise, it returns error
// only exported and addressable fields will be extended
func ExtendStructWithSettableFields(value interface{}) (Builder, error) {
	return MergeStructsWithSettableFields(value)
}

// MergeStructsWithSettableFields merges a list of existing instances of structs and
// returns new instance of Builder interface.
//
// builder := dynamicstruct.MergeStructsWithSettableFields(MyStructOne{}, MyStructTwo{}, MyStructThree{})
//
// it will build only with exported and addressable fields,
// therefore it will not panic, even there are unexported or unaddressable fields in the struct,
// each value in values must be a pointer to struct, otherwise, it returns error
// only exported and addressable fields will be merged
func MergeStructsWithSettableFields(values ...interface{}) (Builder, error) {
	builder := NewStruct()

	for _, value := range values {
		if reflect.TypeOf(value).Kind() != reflect.Ptr {
			return nil, errors.New("values must be pointers to struct")
		}

		elem := reflect.ValueOf(value).Elem()

		for i := 0; i < elem.NumField(); i++ {
			fval := elem.Field(i)
			ftyp := reflect.ValueOf(value).Type().Elem().Field(i)
			ef := elem.Field(i)
			if ef.CanSet() {
				builder.AddField(ftyp.Name, fval.Interface(), string(ftyp.Tag))
			}
		}
	}

	return builder, nil
}

func (b *builderImpl) AddField(name string, typ interface{}, tag string) Builder {
	b.fields[name] = &fieldConfigImpl{
		typ: typ,
		tag: tag,
	}

	return b
}

func (b *builderImpl) RemoveField(name string) Builder {
	delete(b.fields, name)

	return b
}

func (b *builderImpl) HasField(name string) bool {
	_, ok := b.fields[name]
	return ok
}

func (b *builderImpl) GetField(name string) FieldConfig {
	if !b.HasField(name) {
		return nil
	}
	return b.fields[name]
}

func (b *builderImpl) Build() DynamicStruct {
	var structFields []reflect.StructField

	for name, field := range b.fields {
		structFields = append(structFields, reflect.StructField{
			Name: name,
			Type: reflect.TypeOf(field.typ),
			Tag:  reflect.StructTag(field.tag),
		})
	}

	return &dynamicStructImpl{
		definition: reflect.StructOf(structFields),
	}
}

func (f *fieldConfigImpl) SetType(typ interface{}) FieldConfig {
	f.typ = typ
	return f
}

func (f *fieldConfigImpl) SetTag(tag string) FieldConfig {
	f.tag = tag
	return f
}

func (ds *dynamicStructImpl) New() interface{} {
	return reflect.New(ds.definition).Interface()
}

func (ds *dynamicStructImpl) NewSliceOfStructs() interface{} {
	return reflect.New(reflect.SliceOf(ds.definition)).Interface()
}

func (ds *dynamicStructImpl) NewMapOfStructs(key interface{}) interface{} {
	return reflect.New(reflect.MapOf(reflect.Indirect(reflect.ValueOf(key)).Type(), ds.definition)).Interface()
}
