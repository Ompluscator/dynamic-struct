package dynamicstruct

import (
	"encoding/json"
	"errors"
	"reflect"
)

type (
	Reader interface {
		HasField(name string) bool
		GetField(name string) FieldReader
		ToStruct(out interface{}) error
	}

	FieldReader interface {
		NilInt() *int
		Int() int
		NilInt8() *int8
		Int8() int8
		NilInt16() *int16
		Int16() int16
		NilInt32() *int32
		Int32() int32
		NilInt64() *int64
		Int64() int64
		NilUint() *uint
		Uint() uint
		NilUint8() *uint8
		Uint8() uint8
		NilUint16() *uint16
		Uint16() uint16
		NilUint32() *uint32
		Uint32() uint32
		NilUint64() *uint64
		Uint64() uint64
		NilFloat32() *float32
		Float32() float32
		NilFloat64() *float64
		Float64() float64
		NilString() *string
		String() string
		NilBool() *bool
		Bool() bool
		Interface() interface{}
		MapTo(out interface{}) error
	}

	reader struct {
		value  interface{}
		fields map[string]fieldReader
	}

	fieldReader struct {
		reflect.Value
	}
)

func NewReader(value interface{}) Reader {
	fields := map[string]fieldReader{}

	valueOf := reflect.Indirect(reflect.ValueOf(value))
	typeOf := valueOf.Type()

	for i := 0; i < valueOf.NumField(); i++ {
		fval := valueOf.Field(i)
		ftyp := typeOf.Field(i)
		fields[ftyp.Name] = fieldReader{
			fval,
		}
	}

	return reader{
		value:  value,
		fields: fields,
	}
}

func (r reader) HasField(name string) bool {
	_, ok := r.fields[name]
	return ok
}

func (r reader) GetField(name string) FieldReader {
	if !r.HasField(name) {
		return nil
	}
	return r.fields[name]
}

func (r reader) ToStruct(out interface{}) error {
	valueOf := reflect.ValueOf(out)

	if valueOf.Kind() != reflect.Ptr {
		return errors.New("MapToStruct: expect pointer to be passed")
	}

	return mapStructFields(reflect.Indirect(reflect.ValueOf(r.value)), valueOf)
}

func (f fieldReader) NilInt() *int {
	if f.IsNil() {
		return nil
	}
	value := f.Int()
	return &value
}

func (f fieldReader) Int() int {
	return int(reflect.Indirect(f.Value).Int())
}

func (f fieldReader) NilInt8() *int8 {
	if f.IsNil() {
		return nil
	}
	value := f.Int8()
	return &value
}

func (f fieldReader) Int8() int8 {
	return int8(reflect.Indirect(f.Value).Int())
}

func (f fieldReader) NilInt16() *int16 {
	if f.IsNil() {
		return nil
	}
	value := f.Int16()
	return &value
}

func (f fieldReader) Int16() int16 {
	return int16(reflect.Indirect(f.Value).Int())
}

func (f fieldReader) NilInt32() *int32 {
	if f.IsNil() {
		return nil
	}
	value := f.Int32()
	return &value
}

func (f fieldReader) Int32() int32 {
	return int32(reflect.Indirect(f.Value).Int())
}

func (f fieldReader) NilInt64() *int64 {
	if f.IsNil() {
		return nil
	}
	value := f.Int64()
	return &value
}

func (f fieldReader) Int64() int64 {
	return reflect.Indirect(f.Value).Int()
}

func (f fieldReader) NilUint() *uint {
	if f.IsNil() {
		return nil
	}
	value := f.Uint()
	return &value
}

func (f fieldReader) Uint() uint {
	return uint(reflect.Indirect(f.Value).Uint())
}

func (f fieldReader) NilUint8() *uint8 {
	if f.IsNil() {
		return nil
	}
	value := f.Uint8()
	return &value
}

func (f fieldReader) Uint8() uint8 {
	return uint8(reflect.Indirect(f.Value).Uint())
}

func (f fieldReader) NilUint16() *uint16 {
	if f.IsNil() {
		return nil
	}
	value := f.Uint16()
	return &value
}

func (f fieldReader) Uint16() uint16 {
	return uint16(reflect.Indirect(f.Value).Uint())
}

func (f fieldReader) NilUint32() *uint32 {
	if f.IsNil() {
		return nil
	}
	value := f.Uint32()
	return &value
}

func (f fieldReader) Uint32() uint32 {
	return uint32(reflect.Indirect(f.Value).Uint())
}

func (f fieldReader) NilUint64() *uint64 {
	if f.IsNil() {
		return nil
	}
	value := f.Uint64()
	return &value
}

func (f fieldReader) Uint64() uint64 {
	return reflect.Indirect(f.Value).Uint()
}

func (f fieldReader) NilFloat32() *float32 {
	if f.IsNil() {
		return nil
	}
	value := f.Float32()
	return &value
}

func (f fieldReader) Float32() float32 {
	return float32(reflect.Indirect(f.Value).Float())
}

func (f fieldReader) NilFloat64() *float64 {
	if f.IsNil() {
		return nil
	}
	value := f.Float64()
	return &value
}

func (f fieldReader) Float64() float64 {
	return reflect.Indirect(f.Value).Float()
}

func (f fieldReader) NilBool() *bool {
	if f.IsNil() {
		return nil
	}
	value := f.Bool()
	return &value
}

func (f fieldReader) Bool() bool {
	return reflect.Indirect(f.Value).Bool()
}

func (f fieldReader) NilString() *string {
	if f.IsNil() {
		return nil
	}
	value := f.String()
	return &value
}

func (f fieldReader) String() string {
	return reflect.Indirect(f.Value).String()
}

func (f fieldReader) ReflectValue() reflect.Value {
	return f.Value
}

func (f fieldReader) MapTo(out interface{}) error {
	data, err := json.Marshal(f.Value.Interface())
	if err != nil {
		return err
	}

	return json.Unmarshal(data, out)
}

