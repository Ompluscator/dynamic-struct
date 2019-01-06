package dynamicstruct

import (
	"reflect"
	"testing"
	"time"
)

type (
	testStructOne struct {
		String   string
		Integer  int
		Float    float64
		Bool     bool
		Time     time.Time
		Integers []int
		Custom   testSubStructOne
	}

	testSubStructOne struct {
		String   string
		Integer  int
		Float    float64
		Bool     bool
		Time     time.Time
		Integers []int
	}

	testStructTwo struct {
		String   *string
		Integer  *int
		Float    *float64
		Bool     *bool
		Time     *time.Time
		Integers *[]int
		Custom   *testSubStructTwo
	}

	testSubStructTwo struct {
		String   *string
		Integer  *int
		Float    *float64
		Bool     *bool
		Time     *time.Time
		Integers *[]int
	}
)

func TestReader_GetField(t *testing.T) {
	reader := NewReader(testStructOne{
		String: "some text",
	})

	if _, ok := reader.GetField("String").(fieldReader); !ok {
		t.Error(`TestReader_GetField - expected to have field "String"`)
	}
	if _, ok := reader.GetField("Unknown").(fieldReader); ok {
		t.Error(`TestReader_GetField - expected not to have field "Unknown"`)
	}
}

func TestReader_HasField(t *testing.T) {
	reader := NewReader(testStructOne{
		String: "some text",
	})

	if !reader.HasField("String") {
		t.Error(`TestReader_HasField - expected to have field "String"`)
	}
	if reader.HasField("Unknown") {
		t.Error(`TestReader_HasField - expected not to have field "Unknown"`)
	}
}

func TestReader_ToStruct_FullValueToValue(t *testing.T) {
	text := "some text"
	integer := 123
	float := 123.45
	boolean := true
	yesterday := time.Now().Add(-24 * time.Hour)
	integers := []int{1, 2, 3}

	reader := NewReader(testStructOne{
		String:   text,
		Integer:  integer,
		Float:    float,
		Bool:     boolean,
		Time:     yesterday,
		Integers: integers,
		Custom: testSubStructOne{
			String:   text,
			Integer:  integer,
			Float:    float,
			Bool:     boolean,
			Time:     yesterday,
			Integers: integers,
		},
	})

	expected := testStructOne{
		String:   text,
		Integer:  integer,
		Float:    float,
		Bool:     boolean,
		Time:     yesterday,
		Integers: integers,
		Custom: testSubStructOne{
			String:   text,
			Integer:  integer,
			Float:    float,
			Bool:     boolean,
			Time:     yesterday,
			Integers: integers,
		},
	}

	value := testStructOne{}
	err := reader.ToStruct(&value)
	if err != nil {
		t.Errorf(`TestReader_ToStruct_FullValueToValue - expected not to have error got %#v`, err)
	}

	if !reflect.DeepEqual(expected, value) {
		t.Errorf(`TestReader_ToStruct_FullValueToValue - expected mapped instance to be %#v got %#v`, expected, value)
	}
}

func TestReader_ToStruct_EmptyValueToValue(t *testing.T) {
	reader := NewReader(testStructOne{})

	expected := testStructOne{}

	value := testStructOne{}
	err := reader.ToStruct(&value)
	if err != nil {
		t.Errorf(`TestReader_ToStruct_EmptyValueToValue - expected not to have error got %#v`, err)
	}

	if !reflect.DeepEqual(expected, value) {
		t.Errorf(`TestReader_ToStruct_EmptyValueToValue - expected mapped instance to be %#v got %#v`, expected, value)
	}
}

func TestReader_ToStruct_FullValueToPointer(t *testing.T) {
	text := "some text"
	integer := 123
	float := 123.45
	boolean := true
	yesterday := time.Now().Add(-24 * time.Hour)
	integers := []int{1, 2, 3}

	reader := NewReader(testStructOne{
		String:   text,
		Integer:  integer,
		Float:    float,
		Bool:     boolean,
		Time:     yesterday,
		Integers: integers,
		Custom: testSubStructOne{
			String:   text,
			Integer:  integer,
			Float:    float,
			Bool:     boolean,
			Time:     yesterday,
			Integers: integers,
		},
	})

	expected := testStructTwo{
		String:   &text,
		Integer:  &integer,
		Float:    &float,
		Bool:     &boolean,
		Time:     &yesterday,
		Integers: &integers,
		Custom: &testSubStructTwo{
			String:   &text,
			Integer:  &integer,
			Float:    &float,
			Bool:     &boolean,
			Time:     &yesterday,
			Integers: &integers,
		},
	}

	value := testStructTwo{}
	err := reader.ToStruct(&value)
	if err != nil {
		t.Errorf(`TestReader_ToStruct_FullValueToPointer - expected not to have error got %#v`, err)
	}

	if !reflect.DeepEqual(expected, value) {
		t.Errorf(`TestReader_ToStruct_FullValueToPointer - expected mapped instance to be %#v got %#v`, expected, value)
	}
}

func TestReader_ToStruct_EmptyValueToPointer(t *testing.T) {
	reader := NewReader(testStructOne{})

	text := ""
	integer := 0
	float := 0.0
	boolean := false
	yesterday := time.Time{}
	var integers []int

	expected := testStructTwo{
		String:   &text,
		Integer:  &integer,
		Float:    &float,
		Bool:     &boolean,
		Time:     &yesterday,
		Integers: &integers,
		Custom: &testSubStructTwo{
			String:   &text,
			Integer:  &integer,
			Float:    &float,
			Bool:     &boolean,
			Time:     &yesterday,
			Integers: &integers,
		},
	}

	value := testStructTwo{}
	err := reader.ToStruct(&value)
	if err != nil {
		t.Errorf(`TestReader_ToStruct_EmptyValueToPointer - expected not to have error got %#v`, err)
	}

	if !reflect.DeepEqual(expected, value) {
		t.Errorf(`TestReader_ToStruct_EmptyValueToPointer - expected mapped instance to be %#v got %#v`, expected, value)
	}
}

func TestReader_ToStruct_FullPointerToValue(t *testing.T) {
	text := "some text"
	integer := 123
	float := 123.45
	boolean := true
	yesterday := time.Now().Add(-24 * time.Hour)
	integers := []int{1, 2, 3}

	reader := NewReader(testStructTwo{
		String:   &text,
		Integer:  &integer,
		Float:    &float,
		Bool:     &boolean,
		Time:     &yesterday,
		Integers: &integers,
		Custom: &testSubStructTwo{
			String:   &text,
			Integer:  &integer,
			Float:    &float,
			Bool:     &boolean,
			Time:     &yesterday,
			Integers: &integers,
		},
	})

	expected := testStructOne{
		String:   text,
		Integer:  integer,
		Float:    float,
		Bool:     boolean,
		Time:     yesterday,
		Integers: integers,
		Custom: testSubStructOne{
			String:   text,
			Integer:  integer,
			Float:    float,
			Bool:     boolean,
			Time:     yesterday,
			Integers: integers,
		},
	}

	value := testStructOne{}
	err := reader.ToStruct(&value)
	if err != nil {
		t.Errorf(`TestReader_ToStruct_FullPointerToValue - expected not to have error got %#v`, err)
	}

	if !reflect.DeepEqual(expected, value) {
		t.Errorf(`TestReader_ToStruct_FullPointerToValue - expected mapped instance to be %#v got %#v`, expected, value)
	}
}

func TestReader_ToStruct_EmptyPointerToValue(t *testing.T) {
	reader := NewReader(testStructTwo{})

	expected := testStructOne{}

	value := testStructOne{}
	err := reader.ToStruct(&value)
	if err != nil {
		t.Errorf(`TestReader_ToStruct_EmptyPointerToValue - expected not to have error got %#v`, err)
	}

	if !reflect.DeepEqual(expected, value) {
		t.Errorf(`TestReader_ToStruct_EmptyPointerToValue - expected mapped instance to be %#v got %#v`, expected, value)
	}
}

func TestReader_ToStruct_FullPointerToPointer(t *testing.T) {
	text := "some text"
	integer := 123
	float := 123.45
	boolean := true
	yesterday := time.Now().Add(-24 * time.Hour)
	integers := []int{1, 2, 3}

	reader := NewReader(testStructTwo{
		String:   &text,
		Integer:  &integer,
		Float:    &float,
		Bool:     &boolean,
		Time:     &yesterday,
		Integers: &integers,
		Custom: &testSubStructTwo{
			String:   &text,
			Integer:  &integer,
			Float:    &float,
			Bool:     &boolean,
			Time:     &yesterday,
			Integers: &integers,
		},
	})

	expected := testStructTwo{
		String:   &text,
		Integer:  &integer,
		Float:    &float,
		Bool:     &boolean,
		Time:     &yesterday,
		Integers: &integers,
		Custom: &testSubStructTwo{
			String:   &text,
			Integer:  &integer,
			Float:    &float,
			Bool:     &boolean,
			Time:     &yesterday,
			Integers: &integers,
		},
	}

	value := testStructTwo{}
	err := reader.ToStruct(&value)
	if err != nil {
		t.Errorf(`TestReader_ToStruct_FullPointerToPointer - expected not to have error got %#v`, err)
	}

	if !reflect.DeepEqual(expected, value) {
		t.Errorf(`TestReader_ToStruct_FullPointerToPointer - expected mapped instance to be %#v got %#v`, expected, value)
	}
}

func TestReader_ToStruct_EmptyPointerToPointer(t *testing.T) {
	reader := NewReader(testStructTwo{})

	expected := testStructTwo{}

	value := testStructTwo{}
	err := reader.ToStruct(&value)
	if err != nil {
		t.Errorf(`TestReader_ToStruct_EmptyPointerToPointer - expected not to have error got %#v`, err)
	}

	if !reflect.DeepEqual(expected, value) {
		t.Errorf(`TestReader_ToStruct_EmptyPointerToPointer - expected mapped instance to be %#v got %#v`, expected, value)
	}
}
