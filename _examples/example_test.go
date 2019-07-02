package _examples

import (
	"reflect"
	"testing"
	"time"

	"github.com/ompluscator/dynamic-struct"
	"gopkg.in/go-playground/validator.v9"
)

func TestExample(t *testing.T) {
	instances := []interface{}{
		getReaderWithNewStructForJsonExample(),
		getReaderWithExtendedStructForJsonExample(),
		getReaderWithMergedStructsForJsonExample(),
		getReaderWithNewStructForFormExample(),
		getReaderWithExtendedStructForFormExample(),
		getReaderWithMergedStructsForFormExample(),
	}

	for _, instance := range instances {
		if instance == nil {
			t.Error(`TestExample - expected not to get nil`)
		}

		reader := dynamicstruct.NewReader(instance)

		testInstance(reader, t)
	}

	slices := []interface{}{
		getSliceOfReadersWithNewStructForJsonExample(),
		getSliceOfReadersWithExtendedStructForJsonExample(),
		getSliceOfReadersWithMergedStructsForJsonExample(),
	}

	for _, value := range slices {
		if value == nil {
			t.Error(`TestExample - expected not to get nil`)
		}

		readers := dynamicstruct.NewReader(value).ToSliceOfReaders()

		for _, reader := range readers {
			testInstance(reader, t)
		}
	}

	maps := []interface{}{
		getMapOfReadersWithNewStructForJsonExample(),
		getMapOfReadersWithExtendedStructForJsonExample(),
		getMapOfReadersWithMergedStructsForJsonExample(),
	}

	for _, value := range maps {
		if value == nil {
			t.Error(`TestExample - expected not to get nil`)
		}

		readers := dynamicstruct.NewReader(value).ToSliceOfReaders()

		for _, reader := range readers {
			testInstance(reader, t)
		}
	}
}

func testInstance(reader dynamicstruct.Reader, t *testing.T) {
	if value := reader.GetField("Integer").Int(); value != 123 {
		t.Errorf(`TestExample - expected field "Integer" to be %#v got %#v`, 123, value)
	}

	if value := reader.GetField("Uinteger").Uint(); value != uint(456) {
		t.Errorf(`TestExample - expected field "Uinteger" to be %#v got %#v`, uint(456), value)
	}

	if value := reader.GetField("Text").String(); value != "example" {
		t.Errorf(`TestExample - expected field "Text" to be %#v got %#v`, "example", value)
	}

	if value := reader.GetField("Float").Float64(); value != 123.45 {
		t.Errorf(`TestExample - expected field "Float" to be %#v got %#v`, 123.45, value)
	}

	dateTime, err := time.Parse(time.RFC3339, "2018-12-27T19:42:31+07:00")
	if err != nil {
		t.Errorf(`TestExample - expected not to get error got %#v`, err)
	}
	if value := reader.GetField("Time").Time(); !reflect.DeepEqual(value, dateTime) {
		t.Errorf(`TestExample - expected field "Time" to be %#v got %#v`, dateTime, value)
	}

	if value, ok := reader.GetField("Slice").Interface().([]int); !ok || !reflect.DeepEqual(value, []int{1, 2, 3}) {
		t.Errorf(`TestExample - expected field "Slice" to be %#v got %#v`, []int{1, 2, 3}, value)
	}

	if value := reader.GetField("PointerInteger").PointerInt(); *value != 345 {
		t.Errorf(`TestExample - expected field "PointerInteger" to be %#v got %#v`, 345, *value)
	}

	if value := reader.GetField("PointerUinteger").PointerUint(); *value != uint(234) {
		t.Errorf(`TestExample - expected field "PointerUinteger" to be %#v got %#v`, uint(234), *value)
	}

	if value := reader.GetField("PointerFloat").PointerFloat64(); *value != 567.89 {
		t.Errorf(`TestExample - expected field "PointerFloat" to be %#v got %#v`, 567.89, *value)
	}

	if value := reader.GetField("PointerText").PointerString(); *value != "pointer example" {
		t.Errorf(`TestExample - expected field "PointerText" to be %#v got %#v`, "pointer example", *value)
	}

	if value := reader.GetField("PointerBoolean").PointerBool(); *value != true {
		t.Errorf(`TestExample - expected field "PointerBoolean" to be %#v got %#v`, true, *value)
	}

	pointerDateTime, err := time.Parse(time.RFC3339, "2018-12-28T01:23:45+07:00")
	if err != nil {
		t.Errorf(`TestExample - expected not to get error got %#v`, err)
	}
	if value := reader.GetField("PointerTime").PointerTime(); !reflect.DeepEqual(value, &pointerDateTime) {
		t.Errorf(`TestExample - expected field "PointerTime" to be %#v got %#v`, pointerDateTime, *value)
	}

	if value := reader.GetField("Anonymous").String(); value != "" {
		t.Errorf(`TestExample - expected field "Anonymous" to be empty got %#v`, value)
	}

	subReader := dynamicstruct.NewReader(reader.GetField("SubStruct").Interface())

	if value := subReader.GetField("Integer").Int(); value != 12 {
		t.Errorf(`TestExample - expected field "Integer" to be %#v got %#v`, 12, value)
	}

	if value := subReader.GetField("Text").String(); value != "sub example" {
		t.Errorf(`TestExample - expected field "Text" to be %#v got %#v`, "sub example", value)
	}

	err = validator.New().Struct(reader.GetValue())
	if err == nil {
		t.Errorf(`TestExample - expected to have error got %#v`, err)
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		t.Errorf(`TestExample - expected instance of *validator.ValidationErrors got %#v`, reader.GetValue())
	}

	for _, fieldError := range validationErrors {
		fieldError.Tag()
	}

	fieldErrors := []validator.FieldError(validationErrors)
	if len(fieldErrors) != 2 {
		t.Errorf(`TestExample - expected field errors to have length %#v got %#v`, 2, len(fieldErrors))
	}

	checkedInteger := false
	checkedText := false
	for _, fieldError := range fieldErrors {
		if fieldError.Field() == "Integer" {
			checkedInteger = true

			if fieldError.Tag() != "lt" {
				t.Errorf(`TestExample - expected tag of field error to be %#v got %#v`, "lt", fieldError.Tag())
			}

			if fieldError.Param() != "123" {
				t.Errorf(`TestExample - expected param of field error to be %#v got %#v`, "123", fieldError.Param())
			}
		} else if fieldError.Field() == "Anonymous" {
			checkedText = true

			if fieldError.Tag() != "required" {
				t.Errorf(`TestExample - expected tag of field error to be %#v got %#v`, "required", fieldError.Tag())
			}
		}
	}

	if !checkedInteger {
		t.Error(`TestExample - expected to have field errors for field "Integer"`)
	}

	if !checkedText {
		t.Error(`TestExample - expected to have field errors for field "Anonymous"`)
	}
}
