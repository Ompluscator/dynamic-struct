package dynamic_struct

import "reflect"

type (
	Builder interface {
		AddField(name string, typ interface{}, tag string) Builder
		RemoveField(name string) Builder
		HasField(name string) bool
		GetField(name string) Field
		Build() interface{}
	}

	Field interface {
		SetType(typ interface{}) Field
		SetTag(tag string) Field
	}

	Reader interface {
		Int(name string) int
		Int8(name string) int8
		Int16(name string) int16
		Int32(name string) int32
		Int64(name string) int64
		Uint(name string) uint
		Uint8(name string) uint8
		Uint16(name string) uint16
		Uint32(name string) uint32
		Uint64(name string) uint64
		Float32(name string) float32
		Float64(name string) float64
		String(name string) string
		Bool(name string) bool
	}

	field struct {
		typ interface{}
		tag string
	}

	builder struct {
		fields map[string]*field
	}

	reader struct {
		fields map[string]reflect.Value
	}
)

func NewBuilder() Builder {
	return &builder{
		fields: map[string]*field{},
	}
}

func (b *builder) AddField(name string, typ interface{}, tag string) Builder {
	b.fields[name] = &field{
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

func (b *builder) GetField(name string) Field {
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

func (f *field) SetType(typ interface{}) Field {
	f.typ = typ
	return f
}

func (f *field) SetTag(tag string) Field {
	f.tag = tag
	return f
}

func NewReader(value interface{}) Reader {
	fields := map[string]reflect.Value{}

	valueOf := reflect.ValueOf(value)
	typeOf := reflect.TypeOf(value)

	if typeOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
		typeOf = typeOf.Elem()
	}

	for i := 0; i < valueOf.NumField(); i++ {
		fval := valueOf.Field(i)
		ftyp := typeOf.Field(i)
		fields[ftyp.Name] = fval
	}

	return &reader{
		fields: fields,
	}
}

func (f *reader) Int(name string) int {
	return int(f.fields[name].Int())
}

func (f *reader) Int8(name string) int8 {
	return int8(f.fields[name].Int())
}

func (f *reader) Int16(name string) int16 {
	return int16(f.fields[name].Int())
}

func (f *reader) Int32(name string) int32 {
	return int32(f.fields[name].Int())
}

func (f *reader) Int64(name string) int64 {
	return f.fields[name].Int()
}

func (f *reader) Uint(name string) uint {
	return uint(f.fields[name].Uint())
}

func (f *reader) Uint8(name string) uint8 {
	return uint8(f.fields[name].Uint())
}

func (f *reader) Uint16(name string) uint16 {
	return uint16(f.fields[name].Uint())
}

func (f *reader) Uint32(name string) uint32 {
	return uint32(f.fields[name].Uint())
}

func (f *reader) Uint64(name string) uint64 {
	return f.fields[name].Uint()
}

func (f *reader) Bool(name string) bool {
	return f.fields[name].Bool()
}

func (f *reader) String(name string) string {
	return f.fields[name].String()
}

func (f *reader) Float32(name string) float32 {
	return float32(f.fields[name].Float())
}

func (f *reader) Float64(name string) float64 {
	return f.fields[name].Float()
}
