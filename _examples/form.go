package _examples

import (
	"net/url"
	"time"

	"github.com/go-playground/form"
	"github.com/leebenson/conform"
	"github.com/ompluscator/dynamic-struct"
)

func getReaderWithNewStructForFormExample() interface{} {
	integer := 0
	uinteger := uint(0)
	str := ""
	float := 0.0
	boolean := false

	subInstance := dynamicstruct.NewStruct().
		AddField("Integer", integer, "").
		AddField("Text", str, `form:"subText" conform:"trim"`).
		Build().
		New()

	instance := dynamicstruct.NewStruct().
		AddField("Integer", integer, `form:"int" validate:"lt=123"`).
		AddField("Uinteger", uinteger, `validate:"gte=0"`).
		AddField("Text", str, `form:"someText" conform:"trim"`).
		AddField("Float", float, `form:"double"`).
		AddField("Boolean", boolean, "").
		AddField("Time", time.Time{}, "").
		AddField("Slice", []int{}, "").
		AddField("PointerInteger", &integer, "").
		AddField("PointerUinteger", &uinteger, "").
		AddField("PointerText", &str, `conform:"trim"`).
		AddField("PointerFloat", &float, "").
		AddField("PointerBoolean", &boolean, "").
		AddField("PointerTime", &time.Time{}, "").
		AddField("SubStruct", subInstance, `form:"subData"`).
		AddField("Anonymous", "", `form:"-" validate:"required"`).
		Build().
		New()

	data := url.Values{
		"int":             []string{"123"},
		"Uinteger":        []string{"456"},
		"someText":        []string{" example "},
		"double":          []string{"123.45"},
		"Boolean":         []string{"true"},
		"Time":            []string{"2018-12-27T19:42:31+07:00"},
		"Slice":           []string{"1", "2", "3"},
		"PointerInteger":  []string{"345"},
		"PointerUinteger": []string{"234"},
		"PointerFloat":    []string{"567.89"},
		"PointerText":     []string{" pointer example "},
		"PointerBoolean":  []string{"true"},
		"PointerTime":     []string{"2018-12-28T01:23:45+07:00"},
		"subData.Integer": []string{"12"},
		"subData.subText": []string{" sub example "},
		"Anonymous":       []string{"avoid to read"},
	}

	decoder := form.NewDecoder()

	err := decoder.Decode(&instance, data)
	if err != nil {
		return nil
	}

	err = conform.Strings(instance)
	if err != nil {
		return nil
	}

	return instance
}

func getReaderWithExtendedStructForFormExample() interface{} {
	integer := 0
	uinteger := uint(0)
	str := ""
	float := 0.0
	boolean := false

	instance := dynamicstruct.ExtendStruct(struct {
		Integer   int `form:"int" validate:"lt=123"`
		Uinteger  uint `validate:"gte=0"`
		Text      string  `form:"someText" conform:"trim"`
		Float     float64 `form:"double"`
		Boolean   bool
		Slice     []int
		Time      time.Time
		SubStruct struct {
			Integer int
			Text    string `form:"subText" conform:"trim"`
		} `form:"subData"`
	}{}).
		AddField("PointerInteger", &integer, "").
		AddField("PointerUinteger", &uinteger, "").
		AddField("PointerText", &str, `conform:"trim"`).
		AddField("PointerFloat", &float, "").
		AddField("PointerBoolean", &boolean, "").
		AddField("PointerTime", &time.Time{}, "").
		AddField("Anonymous", "", `form:"-" validate:"required"`).
		Build().
		New()

	data := url.Values{
		"int":             []string{"123"},
		"Uinteger":        []string{"456"},
		"someText":        []string{" example "},
		"double":          []string{"123.45"},
		"Boolean":         []string{"true"},
		"Time":            []string{"2018-12-27T19:42:31+07:00"},
		"Slice":           []string{"1", "2", "3"},
		"PointerInteger":  []string{"345"},
		"PointerUinteger": []string{"234"},
		"PointerFloat":    []string{"567.89"},
		"PointerText":     []string{" pointer example "},
		"PointerBoolean":  []string{"true"},
		"PointerTime":     []string{"2018-12-28T01:23:45+07:00"},
		"subData.Integer": []string{"12"},
		"subData.subText": []string{" sub example "},
		"Anonymous":       []string{"avoid to read"},
	}

	decoder := form.NewDecoder()

	err := decoder.Decode(&instance, data)
	if err != nil {
		return nil
	}

	err = conform.Strings(instance)
	if err != nil {
		return nil
	}

	return instance
}

func getReaderWithMergedStructsForFormExample() interface{} {
	instance := dynamicstruct.MergeStructs(struct {
		Integer  int `form:"int" validate:"lt=123"`
		Uinteger uint `validate:"gte=0"`
		Text     string  `form:"someText" conform:"trim"`
		Float    float64 `form:"double"`
		Boolean  bool
		Slice    []int
		Time     time.Time
	}{}, struct {
		Anonymous string `form:"-" validate:"required"`
		SubStruct struct {
			Integer int
			Text    string `form:"subText" conform:"trim"`
		} `form:"subData"`
	}{}, struct {
		PointerInteger  *int
		PointerUinteger *uint
		PointerText     *string `conform:"trim"`
		PointerFloat    *float64
		PointerBoolean  *bool
		PointerTime     *time.Time
	}{}).
		Build().
		New()

	data := url.Values{
		"int":             []string{"123"},
		"Uinteger":        []string{"456"},
		"someText":        []string{" example "},
		"double":          []string{"123.45"},
		"Boolean":         []string{"true"},
		"Time":            []string{"2018-12-27T19:42:31+07:00"},
		"Slice":           []string{"1", "2", "3"},
		"PointerInteger":  []string{"345"},
		"PointerUinteger": []string{"234"},
		"PointerFloat":    []string{"567.89"},
		"PointerText":     []string{" pointer example "},
		"PointerBoolean":  []string{"true"},
		"PointerTime":     []string{"2018-12-28T01:23:45+07:00"},
		"subData.Integer": []string{"12"},
		"subData.subText": []string{" sub example "},
		"Anonymous":       []string{"avoid to read"},
	}

	decoder := form.NewDecoder()

	err := decoder.Decode(&instance, data)
	if err != nil {
		return nil
	}

	err = conform.Strings(instance)
	if err != nil {
		return nil
	}

	return instance
}
