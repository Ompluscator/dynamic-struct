package dynamic_struct

import "reflect"

type (
	Builder interface {
		AddField(name string, typ interface{}, tag string) Builder
		RemoveField(name string) Builder
		HasField(name string) bool
		GetField(name string) FieldConfig
		Build() interface{}
	}

	FieldConfig interface {
		SetType(typ interface{}) FieldConfig
		SetTag(tag string) FieldConfig
	}

	builder struct {
		fields map[string]*fieldConfig
	}

	fieldConfig struct {
		typ interface{}
		tag string
	}
)

func NewBuilder() Builder {
	return &builder{
		fields: map[string]*fieldConfig{},
	}
}

func (b *builder) AddField(name string, typ interface{}, tag string) Builder {
	b.fields[name] = &fieldConfig{
		typ: typ,
		tag: tag,
	}

	return b
}

func (b *builder) RemoveField(name string) Builder {
	delete(b.fields, name)

	return b
}

func (b *builder) HasField(name string) bool {
	_, ok := b.fields[name]
	return ok
}

func (b *builder) GetField(name string) FieldConfig {
	return b.fields[name]
}

func (b *builder) Build() interface{} {
	var structFields []reflect.StructField

	for name, field := range b.fields {
		structFields = append(structFields, reflect.StructField{
			Name: name,
			Type: reflect.TypeOf(field.typ),
			Tag:  reflect.StructTag(field.tag),
		})
	}

	return reflect.New(reflect.StructOf(structFields)).Interface()
}

func (f *fieldConfig) SetType(typ interface{}) FieldConfig {
	f.typ = typ
	return f
}

func (f *fieldConfig) SetTag(tag string) FieldConfig {
	f.tag = tag
	return f
}
