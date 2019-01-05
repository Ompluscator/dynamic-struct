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

	builderImpl struct {
		fields map[string]*fieldConfigImpl
	}

	fieldConfigImpl struct {
		typ interface{}
		tag string
	}
)

func NewBuilder() Builder {
	return &builderImpl{
		fields: map[string]*fieldConfigImpl{},
	}
}

func ExtendStruct(value interface{}) Builder {
	fields := map[string]*fieldConfigImpl{}

	valueOf := reflect.Indirect(reflect.ValueOf(value))
	typeOf := valueOf.Type()

	for i := 0; i < valueOf.NumField(); i++ {
		fval := valueOf.Field(i)
		ftyp := typeOf.Field(i)
		fields[ftyp.Name] = &fieldConfigImpl{
			typ: fval.Interface(),
			tag: string(ftyp.Tag),
		}
	}

	return &builderImpl{
		fields: fields,
	}
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

func (b *builderImpl) Build() interface{} {
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

func (f *fieldConfigImpl) SetType(typ interface{}) FieldConfig {
	f.typ = typ
	return f
}

func (f *fieldConfigImpl) SetTag(tag string) FieldConfig {
	f.tag = tag
	return f
}
