package dynamicstruct

import (
	"errors"
	"fmt"
	"reflect"
)

func mapStructFields(sourceStruct reflect.Value, destinationStruct reflect.Value) error {
	destinationStruct = reflect.Indirect(destinationStruct)

	if !isStruct(destinationStruct.Type()) {
		return errors.New("MapToStruct: expect pointer to struct to be passed")
	}

	for i := 0; i < destinationStruct.NumField(); i++ {
		mapField(i, destinationStruct, sourceStruct)
	}

	return nil
}

func mapField(i int, destinationStruct reflect.Value, sourceStruct reflect.Value) error {
	destinationStructFieldValue := destinationStruct.Field(i)
	destinationStructField := destinationStruct.Type().Field(i)
	if !isPossibleToSetField(i, sourceStruct, destinationStruct) {
		return nil
	}

	sourceValue := sourceStruct.FieldByName(destinationStructField.Name)

	originalDestinationType := destinationStructFieldValue.Type()
	destinationTrueType := getTrueFieldType(originalDestinationType)

	originalSourceType := sourceValue.Type()
	sourceTrueType := getTrueFieldType(originalSourceType)

	if isStruct(originalDestinationType) && isStruct(originalSourceType) && !areSameStructs(originalDestinationType, originalSourceType) {
		newValue := reflect.Indirect(reflect.New(destinationTrueType))
		if isPointer(originalSourceType) && sourceValue.IsNil() {
			return nil
		}
		err := mapStructFields(reflect.Indirect(sourceValue), newValue)
		if err != nil {
			return err
		}
		setNewValue(destinationStructFieldValue, newValue)
	} else if destinationTrueType.Kind() == sourceTrueType.Kind() {
		newValue := sourceValue
		if !isPointer(originalDestinationType) && isPointer(originalSourceType) && !sourceValue.IsNil() {
			newValue = sourceValue.Elem()
		}
		setNewValue(destinationStructFieldValue, newValue)
	} else if sourceTrueType.ConvertibleTo(destinationTrueType) {
		var newValue reflect.Value
		if !isPointer(originalDestinationType) && isPointer(originalSourceType) && !sourceValue.IsNil() {
			newValue = sourceValue.Elem().Convert(destinationTrueType)
		} else {
			newValue = sourceValue.Convert(destinationTrueType)
		}
		setNewValue(destinationStructFieldValue, newValue)
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

func isStruct(typ reflect.Type) bool {
	return getTrueFieldType(typ).Kind() == reflect.Struct
}

func isPointer(typ reflect.Type) bool {
	return typ.Kind() == reflect.Ptr
}

func isNamedStruct(typ reflect.Type) bool {
	return typ.Name() != "" && typ.PkgPath() != ""
}

func areSameStructs(first reflect.Type, second reflect.Type) bool {
	firstTrue := getTrueFieldType(first)
	secondTrue := getTrueFieldType(second)

	return isNamedStruct(firstTrue) && isNamedStruct(secondTrue) && firstTrue.Name() == secondTrue.Name() && firstTrue.PkgPath() == secondTrue.PkgPath()
}

func setNewValue(destination reflect.Value, source reflect.Value) {
	destinationType := destination.Type()
	sourceType := source.Type()

	if isPointer(destinationType) && !isPointer(sourceType) {
		destination.Set(reflect.New(getTrueFieldType(destinationType)))
		destination.Elem().Set(source)
	} else if !isPointer(destinationType) && isPointer(sourceType) {
		if !source.IsNil() {
			destination.Set(source)
		}
	} else {
		destination.Set(source)
	}
}

func getTrueFieldType(fieldType reflect.Type) reflect.Type {
	if fieldType.Kind() == reflect.Ptr {
		return fieldType.Elem()
	}

	return fieldType
}