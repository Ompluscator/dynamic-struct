package dynamicstruct

import (
	"fmt"
	"reflect"
	"time"
)

type (
	Reader interface {
		HasField(name string) bool
		GetField(name string) Field
	}

	Field interface {
		PointerInt() *int
		Int() int
		PointerInt8() *int8
		Int8() int8
		PointerInt16() *int16
		Int16() int16
		PointerInt32() *int32
		Int32() int32
		PointerInt64() *int64
		Int64() int64
		PointerUint() *uint
		Uint() uint
		PointerUint8() *uint8
		Uint8() uint8
		PointerUint16() *uint16
		Uint16() uint16
		PointerUint32() *uint32
		Uint32() uint32
		PointerUint64() *uint64
		Uint64() uint64
		PointerFloat32() *float32
		Float32() float32
		PointerFloat64() *float64
		Float64() float64
		PointerString() *string
		String() string
		PointerBool() *bool
		Bool() bool
		PointerTime() *time.Time
		Time() time.Time
		Interface() interface{}
	}

	readImpl struct {
		fields map[string]fieldImpl
	}

	fieldImpl struct {
		field reflect.StructField
		value reflect.Value
	}
)

func NewReader(value interface{}) Reader {
	fields := map[string]fieldImpl{}

	valueOf := reflect.Indirect(reflect.ValueOf(value))
	typeOf := valueOf.Type()

	for i := 0; i < valueOf.NumField(); i++ {
		field := typeOf.Field(i)
		fields[field.Name] = fieldImpl{
			field: field,
			value: valueOf.Field(i),
		}
	}

	return readImpl{
		fields: fields,
	}
}

func (r readImpl) HasField(name string) bool {
	_, ok := r.fields[name]
	return ok
}

func (r readImpl) GetField(name string) Field {
	if !r.HasField(name) {
		return nil
	}
	return r.fields[name]
}

func (f fieldImpl) PointerInt() *int {
	if f.value.IsNil() {
		return nil
	}
	value := f.Int()
	return &value
}

func (f fieldImpl) Int() int {
	return int(reflect.Indirect(f.value).Int())
}

func (f fieldImpl) PointerInt8() *int8 {
	if f.value.IsNil() {
		return nil
	}
	value := f.Int8()
	return &value
}

func (f fieldImpl) Int8() int8 {
	return int8(reflect.Indirect(f.value).Int())
}

func (f fieldImpl) PointerInt16() *int16 {
	if f.value.IsNil() {
		return nil
	}
	value := f.Int16()
	return &value
}

func (f fieldImpl) Int16() int16 {
	return int16(reflect.Indirect(f.value).Int())
}

func (f fieldImpl) PointerInt32() *int32 {
	if f.value.IsNil() {
		return nil
	}
	value := f.Int32()
	return &value
}

func (f fieldImpl) Int32() int32 {
	return int32(reflect.Indirect(f.value).Int())
}

func (f fieldImpl) PointerInt64() *int64 {
	if f.value.IsNil() {
		return nil
	}
	value := f.Int64()
	return &value
}

func (f fieldImpl) Int64() int64 {
	return reflect.Indirect(f.value).Int()
}

func (f fieldImpl) PointerUint() *uint {
	if f.value.IsNil() {
		return nil
	}
	value := f.Uint()
	return &value
}

func (f fieldImpl) Uint() uint {
	return uint(reflect.Indirect(f.value).Uint())
}

func (f fieldImpl) PointerUint8() *uint8 {
	if f.value.IsNil() {
		return nil
	}
	value := f.Uint8()
	return &value
}

func (f fieldImpl) Uint8() uint8 {
	return uint8(reflect.Indirect(f.value).Uint())
}

func (f fieldImpl) PointerUint16() *uint16 {
	if f.value.IsNil() {
		return nil
	}
	value := f.Uint16()
	return &value
}

func (f fieldImpl) Uint16() uint16 {
	return uint16(reflect.Indirect(f.value).Uint())
}

func (f fieldImpl) PointerUint32() *uint32 {
	if f.value.IsNil() {
		return nil
	}
	value := f.Uint32()
	return &value
}

func (f fieldImpl) Uint32() uint32 {
	return uint32(reflect.Indirect(f.value).Uint())
}

func (f fieldImpl) PointerUint64() *uint64 {
	if f.value.IsNil() {
		return nil
	}
	value := f.Uint64()
	return &value
}

func (f fieldImpl) Uint64() uint64 {
	return reflect.Indirect(f.value).Uint()
}

func (f fieldImpl) PointerFloat32() *float32 {
	if f.value.IsNil() {
		return nil
	}
	value := f.Float32()
	return &value
}

func (f fieldImpl) Float32() float32 {
	return float32(reflect.Indirect(f.value).Float())
}

func (f fieldImpl) PointerFloat64() *float64 {
	if f.value.IsNil() {
		return nil
	}
	value := f.Float64()
	return &value
}

func (f fieldImpl) Float64() float64 {
	return reflect.Indirect(f.value).Float()
}

func (f fieldImpl) PointerString() *string {
	if f.value.IsNil() {
		return nil
	}
	value := f.String()
	return &value
}

func (f fieldImpl) String() string {
	return reflect.Indirect(f.value).String()
}

func (f fieldImpl) PointerBool() *bool {
	if f.value.IsNil() {
		return nil
	}
	value := f.Bool()
	return &value
}

func (f fieldImpl) Bool() bool {
	return reflect.Indirect(f.value).Bool()
}

func (f fieldImpl) PointerTime() *time.Time {
	if f.value.IsNil() {
		return nil
	}
	value := f.Time()
	return &value
}

func (f fieldImpl) Time() time.Time {
	value, ok := reflect.Indirect(f.value).Interface().(time.Time)
	if !ok {
		panic(fmt.Sprintf(`field "%s" is not instance of time.Time`, f.field.Name))
	}

	return value
}

func (f fieldImpl) Interface() interface{} {
	return f.value.Interface()
}
