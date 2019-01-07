package dynamicstruct

import (
	"errors"
	"fmt"
	"reflect"
)

func mapStructFields(sourceStruct reflect.Value, destinationStruct reflect.Value) error {
	destinationStruct = reflect.Indirect(destinationStruct)

	if !isOriginallyStruct(destinationStruct) {
		return errors.New("MapToStruct: expect pointer to struct to be passed")
	}

	for i := 0; i < destinationStruct.NumField(); i++ {
		err := mapField(i, destinationStruct, sourceStruct)
		if err != nil {
			return err
		}
	}

	return nil
}

func mapField(i int, destinationStruct reflect.Value, sourceStruct reflect.Value) error {
	destinationValue := destinationStruct.Field(i)
	destinationStructField := destinationStruct.Type().Field(i)
	if !isPossibleToSetField(i, sourceStruct, destinationStruct) {
		return nil
	}

	sourceValue := sourceStruct.FieldByName(destinationStructField.Name)
	destinationType := getUnderlyingValueType(destinationValue)
	sourceType := getUnderlyingValueType(sourceValue)

	if isOriginallyStruct(destinationValue) && isOriginallyStruct(sourceValue) && !areSameStructs(destinationValue, sourceValue) {
		newValue := reflect.Indirect(reflect.New(destinationType))
		if isPointer(sourceValue) && sourceValue.IsNil() {
			return nil
		}
		err := mapStructFields(reflect.Indirect(sourceValue), newValue)
		if err != nil {
			return err
		}
		copyValue(destinationValue, newValue)
	} else if destinationType.Kind() == sourceType.Kind() {
		copyValue(destinationValue, sourceValue)
	} else if sourceType.ConvertibleTo(destinationType) {
		convertValue(destinationValue, sourceValue)
	} else {
		return errors.New(fmt.Sprintf(`MapToStruct: field "%s" is not convertible`, destinationStructField.Name))
	}

	return nil
}

func isPossibleToSetField(i int, sourceStruct reflect.Value, destinationStruct reflect.Value) bool {
	typ := destinationStruct.Type()
	value := destinationStruct.Field(i)

	return sourceStruct.FieldByName(typ.Field(i).Name).IsValid() && value.IsValid() && value.CanSet()
}

func isOriginallyStruct(value reflect.Value) bool {
	return getUnderlyingValueType(value).Kind() == reflect.Struct
}

func isPointer(value reflect.Value) bool {
	return value.Type().Kind() == reflect.Ptr
}

func isStructWithPackageAndName(value reflect.Value) bool {
	typ := getUnderlyingValueType(value)

	return typ.Name() != "" && typ.PkgPath() != ""
}

func areSameStructs(first reflect.Value, second reflect.Value) bool {
	if !isStructWithPackageAndName(first) || !isStructWithPackageAndName(second) {
		return false
	}
	
	firstType := getUnderlyingValueType(first)
	secondType := getUnderlyingValueType(second)

	return firstType.Name() == secondType.Name() && firstType.PkgPath() == secondType.PkgPath()
}

func copyValue(destination reflect.Value, source reflect.Value) {
	if isPointer(source) && source.IsNil() {
		return
	}

	destinationType := getUnderlyingValueType(destination)

	if isPointer(destination) {
		destination.Set(reflect.New(destinationType))
		destination = destination.Elem()
	}

	if isPointer(source) {
		source = source.Elem()
	}

	destination.Set(source)
}

func convertValue(destination reflect.Value, source reflect.Value) {
	if isPointer(source) && source.IsNil() {
		return
	}

	destinationType := getUnderlyingValueType(destination)

	if isPointer(destination) {
		destination.Set(reflect.New(destinationType))
		destination = destination.Elem()
	}

	if isPointer(source) {
		source = source.Elem()
	}

	destination.Set(source.Convert(destinationType))
}

func getUnderlyingValueType(fieldType reflect.Value) reflect.Type {
	typ := fieldType.Type()

	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ
}