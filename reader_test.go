package dynamicstruct

import (
	"reflect"
	"testing"
	"time"
)

type (
	testStruct struct {
		String          string
		Integer         int
		Uinteger        uint
		Float           float64
		Bool            bool
		Time            time.Time
		PointerString   *string
		PointerInteger  *int
		PointerUinteger *uint
		PointerFloat    *float64
		PointerBool     *bool
		PointerTime     *time.Time
		Integers        []int
	}
)

func TestReaderImpl_GetField(t *testing.T) {
	reader := NewReader(testStruct{
		String: "some text",
	})

	if _, ok := reader.GetField("String").(fieldImpl); !ok {
		t.Error(`TestReaderImpl_GetField - expected to have field "String"`)
	}
	if _, ok := reader.GetField("Unknown").(fieldImpl); ok {
		t.Error(`TestReaderImpl_GetField - expected not to have field "Unknown"`)
	}
}

func TestReaderImpl_HasField(t *testing.T) {
	reader := NewReader(testStruct{
		String: "some text",
	})

	if !reader.HasField("String") {
		t.Error(`TestReaderImpl_HasField - expected to have field "String"`)
	}
	if reader.HasField("Unknown") {
		t.Error(`TestReaderImpl_HasField - expected not to have field "Unknown"`)
	}
}

func TestReaderImpl_GetAllFields(t *testing.T) {
	reader := NewReader(testStruct{})

	if len(reader.GetAllFields()) != 13 {
		t.Errorf(`TestReaderImpl_GetAllFields - expected to have 10 fields but got %d`, len(reader.GetAllFields()))
	}
}

func TestFieldImpl_Name(t *testing.T) {
	reader := NewReader(testStruct{})

	if reader.GetField("String").Name() != "String" {
		t.Errorf(`TestFieldImpl_Name - expected field name to be "String"  got %s`, reader.GetField("String").Name())
	}
}

func TestFieldImpl_PointerInt(t *testing.T) {
	expected := 123

	reader := NewReader(testStruct{
		PointerInteger: &expected,
	})

	value := reader.GetField("PointerInteger").PointerInt()

	if *value != expected {
		t.Errorf(`TestFieldImpl_PointerInt - expected field "PointerInteger" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerInteger").PointerInt()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerInt - expected field "PointerInteger" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Int(t *testing.T) {
	expected := 123

	reader := NewReader(testStruct{
		Integer: expected,
	})

	value := reader.GetField("Integer").Int()

	if value != expected {
		t.Errorf(`TestFieldImpl_Int - expected field "Integer" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerInt8(t *testing.T) {
	expected := 123

	reader := NewReader(testStruct{
		PointerInteger: &expected,
	})

	value := reader.GetField("PointerInteger").PointerInt8()

	if *value != int8(expected) {
		t.Errorf(`TestFieldImpl_PointerInt8 - expected field "PointerInteger" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerInteger").PointerInt8()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerInt8 - expected field "PointerInteger" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Int8(t *testing.T) {
	expected := 123

	reader := NewReader(testStruct{
		Integer: expected,
	})

	value := reader.GetField("Integer").Int8()

	if value != int8(expected) {
		t.Errorf(`TestFieldImpl_Int8 - expected field "Integer" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerInt16(t *testing.T) {
	expected := 123

	reader := NewReader(testStruct{
		PointerInteger: &expected,
	})

	value := reader.GetField("PointerInteger").PointerInt16()

	if *value != int16(expected) {
		t.Errorf(`TestFieldImpl_PointerInt16 - expected field "PointerInteger" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerInteger").PointerInt16()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerInt16 - expected field "PointerInteger" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Int16(t *testing.T) {
	expected := 123

	reader := NewReader(testStruct{
		Integer: expected,
	})

	value := reader.GetField("Integer").Int16()

	if value != int16(expected) {
		t.Errorf(`TestFieldImpl_Int16 - expected field "Integer" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerInt32(t *testing.T) {
	expected := 123

	reader := NewReader(testStruct{
		PointerInteger: &expected,
	})

	value := reader.GetField("PointerInteger").PointerInt32()

	if *value != int32(expected) {
		t.Errorf(`TestFieldImpl_PointerInt32 - expected field "PointerInteger" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerInteger").PointerInt32()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerInt32 - expected field "PointerInteger" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Int32(t *testing.T) {
	expected := 123

	reader := NewReader(testStruct{
		Integer: expected,
	})

	value := reader.GetField("Integer").Int32()

	if value != int32(expected) {
		t.Errorf(`TestFieldImpl_Int32 - expected field "Integer" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerInt64(t *testing.T) {
	expected := 123

	reader := NewReader(testStruct{
		PointerInteger: &expected,
	})

	value := reader.GetField("PointerInteger").PointerInt64()

	if *value != int64(expected) {
		t.Errorf(`TestFieldImpl_PointerInt64 - expected field "PointerInteger" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerInteger").PointerInt64()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerInt64 - expected field "PointerInteger" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Int64(t *testing.T) {
	expected := 123

	reader := NewReader(testStruct{
		Integer: expected,
	})

	value := reader.GetField("Integer").Int64()

	if value != int64(expected) {
		t.Errorf(`TestFieldImpl_Int64 - expected field "Integer" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerUint(t *testing.T) {
	expected := uint(123)

	reader := NewReader(testStruct{
		PointerUinteger: &expected,
	})

	value := reader.GetField("PointerUinteger").PointerUint()

	if *value != expected {
		t.Errorf(`TestFieldImpl_PointerUint - expected field "PointerUinteger" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerUinteger").PointerUint()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerUint - expected field "PointerUinteger" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Uint(t *testing.T) {
	expected := uint(123)

	reader := NewReader(testStruct{
		Uinteger: expected,
	})

	value := reader.GetField("Uinteger").Uint()

	if value != expected {
		t.Errorf(`TestFieldImpl_Uint - expected field "Uinteger" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerUint8(t *testing.T) {
	expected := uint(123)

	reader := NewReader(testStruct{
		PointerUinteger: &expected,
	})

	value := reader.GetField("PointerUinteger").PointerUint8()

	if *value != uint8(expected) {
		t.Errorf(`TestFieldImpl_PointerUint8 - expected field "PointerUinteger" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerUinteger").PointerUint8()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerUint8 - expected field "PointerUinteger" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Uint8(t *testing.T) {
	expected := uint(123)

	reader := NewReader(testStruct{
		Uinteger: expected,
	})

	value := reader.GetField("Uinteger").Uint8()

	if value != uint8(expected) {
		t.Errorf(`TestFieldImpl_Uint8 - expected field "Uinteger" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerUint16(t *testing.T) {
	expected := uint(123)

	reader := NewReader(testStruct{
		PointerUinteger: &expected,
	})

	value := reader.GetField("PointerUinteger").PointerUint16()

	if *value != uint16(expected) {
		t.Errorf(`TestFieldImpl_PointerUint16 - expected field "PointerUinteger" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerUinteger").PointerUint16()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerUint16 - expected field "PointerUinteger" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Uint16(t *testing.T) {
	expected := uint(123)

	reader := NewReader(testStruct{
		Uinteger: expected,
	})

	value := reader.GetField("Uinteger").Uint16()

	if value != uint16(expected) {
		t.Errorf(`TestFieldImpl_Uint16 - expected field "Uinteger" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerUint32(t *testing.T) {
	expected := uint(123)

	reader := NewReader(testStruct{
		PointerUinteger: &expected,
	})

	value := reader.GetField("PointerUinteger").PointerUint32()

	if *value != uint32(expected) {
		t.Errorf(`TestFieldImpl_PointerUint32 - expected field "PointerUinteger" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerUinteger").PointerUint32()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerUint32 - expected field "PointerUinteger" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Uint32(t *testing.T) {
	expected := uint(123)

	reader := NewReader(testStruct{
		Uinteger: expected,
	})

	value := reader.GetField("Uinteger").Uint32()

	if value != uint32(expected) {
		t.Errorf(`TestFieldImpl_Uint32 - expected field "Uinteger" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerUint64(t *testing.T) {
	expected := uint(123)

	reader := NewReader(testStruct{
		PointerUinteger: &expected,
	})

	value := reader.GetField("PointerUinteger").PointerUint64()

	if *value != uint64(expected) {
		t.Errorf(`TestFieldImpl_PointerUint64 - expected field "PointerUinteger" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerUinteger").PointerUint64()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerUint64 - expected field "PointerUinteger" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Uint64(t *testing.T) {
	expected := uint(123)

	reader := NewReader(testStruct{
		Uinteger: expected,
	})

	value := reader.GetField("Uinteger").Uint64()

	if value != uint64(expected) {
		t.Errorf(`TestFieldImpl_Uint64 - expected field "Uinteger" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerFloat32(t *testing.T) {
	expected := 123.0

	reader := NewReader(testStruct{
		PointerFloat: &expected,
	})

	value := reader.GetField("PointerFloat").PointerFloat32()

	if *value != float32(expected) {
		t.Errorf(`TestFieldImpl_PointerFloat32 - expected field "PointerFloat" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerFloat").PointerFloat32()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerFloat32 - expected field "PointerFloat" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Float32(t *testing.T) {
	expected := 123.0

	reader := NewReader(testStruct{
		Float: expected,
	})

	value := reader.GetField("Float").Float32()

	if value != float32(expected) {
		t.Errorf(`TestFieldImpl_Float32 - expected field "Float" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerFloat64(t *testing.T) {
	expected := 123.0

	reader := NewReader(testStruct{
		PointerFloat: &expected,
	})

	value := reader.GetField("PointerFloat").PointerFloat64()

	if *value != expected {
		t.Errorf(`TestFieldImpl_PointerFloat64 - expected field "PointerFloat" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerFloat").PointerFloat64()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerFloat64 - expected field "PointerFloat" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Float64(t *testing.T) {
	expected := 123.0

	reader := NewReader(testStruct{
		Float: expected,
	})

	value := reader.GetField("Float").Float64()

	if value != expected {
		t.Errorf(`TestFieldImpl_Float64 - expected field "Float" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerString(t *testing.T) {
	expected := "something"

	reader := NewReader(testStruct{
		PointerString: &expected,
	})

	value := reader.GetField("PointerString").PointerString()

	if *value != expected {
		t.Errorf(`TestFieldImpl_PointerString - expected field "PointerString" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerString").PointerString()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerString - expected field "PointerString" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_String(t *testing.T) {
	expected := "something"

	reader := NewReader(testStruct{
		String: expected,
	})

	value := reader.GetField("String").String()

	if value != expected {
		t.Errorf(`TestFieldImpl_String - expected field "String" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerBool(t *testing.T) {
	expected := true

	reader := NewReader(testStruct{
		PointerBool: &expected,
	})

	value := reader.GetField("PointerBool").PointerBool()

	if *value != expected {
		t.Errorf(`TestFieldImpl_PointerBool - expected field "PointerBool" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerBool").PointerBool()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerBool - expected field "PointerBool" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Bool(t *testing.T) {
	expected := true

	reader := NewReader(testStruct{
		Bool: expected,
	})

	value := reader.GetField("Bool").Bool()

	if value != expected {
		t.Errorf(`TestFieldImpl_Bool - expected field "Bool" to be equal %#v but got %#v`, expected, value)
	}
}

func TestFieldImpl_PointerTime(t *testing.T) {
	expected := time.Now()

	reader := NewReader(testStruct{
		PointerTime: &expected,
	})

	value := reader.GetField("PointerTime").PointerTime()

	if *value != expected {
		t.Errorf(`TestFieldImpl_PointerTime - expected field "PointerTime" to be equal %#v but got %#v`, expected, *value)
	}

	reader = NewReader(testStruct{})

	value = reader.GetField("PointerTime").PointerTime()

	if value != nil {
		t.Errorf(`TestFieldImpl_PointerTime - expected field "PointerTime" to be nil but got %#v`, *value)
	}
}

func TestFieldImpl_Time(t *testing.T) {
	expected := time.Now()

	reader := NewReader(testStruct{
		Time: expected,
	})

	value := reader.GetField("Time").Time()

	if value != expected {
		t.Errorf(`TestFieldImpl_Time - expected field "Time" to be equal %#v but got %#v`, expected, value)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestFieldImpl_Time - expected panic by casting int to instance of time.Time{}")
		}
	}()
	value = reader.GetField("Integer").Time()
}

func TestFieldImpl_Interface(t *testing.T) {
	expected := []int{1, 2, 3}

	reader := NewReader(testStruct{
		Integers: expected,
	})

	value, ok := reader.GetField("Integers").Interface().([]int)

	if !ok {
		t.Error(`TestFieldImpl_Interface - expected field "String" to be instance of []int`)
	}

	if !reflect.DeepEqual(value, expected) {
		t.Errorf(`TestFieldImpl_Interface - expected field "String" to be equal %#v but got %#v`, expected, value)
	}
}
