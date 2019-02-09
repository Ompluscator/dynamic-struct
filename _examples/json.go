package _examples

import (
	"encoding/json"
	"time"

	"github.com/ompluscator/dynamic-struct"
)

func getReaderWithNewStructForJsonExample() interface{} {
	integer := 0
	uinteger := uint(0)
	str := ""
	float := 0.0
	boolean := false

	subInstance := dynamicstruct.NewStruct().
		AddField("Integer", integer, "").
		AddField("Text", str, `json:"subText"`).
		Build().
		New()

	instance := dynamicstruct.NewStruct().
		AddField("Integer", integer, `json:"int" validate:"lt=123"`).
		AddField("Uinteger", uinteger, `validate:"gte=0"`).
		AddField("Text", str, `json:"someText"`).
		AddField("Float", float, `json:"double"`).
		AddField("Boolean", boolean, "").
		AddField("Time", time.Time{}, "").
		AddField("Slice", []int{}, "").
		AddField("PointerInteger", &integer, "").
		AddField("PointerUinteger", &uinteger, "").
		AddField("PointerText", &str, "").
		AddField("PointerFloat", &float, "").
		AddField("PointerBoolean", &boolean, "").
		AddField("PointerTime", &time.Time{}, "").
		AddField("SubStruct", subInstance, `json:"subData"`).
		AddField("Anonymous", "", `json:"-" validate:"required"`).
		Build().
		New()

	data := []byte(`
{
	"int": 123,
	"Uinteger": 456,
	"someText": "example",
	"double": 123.45,
	"Boolean": true,
	"Time": "2018-12-27T19:42:31+07:00",
	"Slice": [1, 2, 3],
	"PointerInteger": 345,
	"PointerUinteger": 234,
	"PointerFloat": 567.89,
	"PointerText": "pointer example",
	"PointerBoolean": true,
	"PointerTime": "2018-12-28T01:23:45+07:00",
	"subData": {
		"Integer": 12,
		"subText": "sub example"
	},
	"Anonymous": "avoid to read"
}
`)

	err := json.Unmarshal(data, &instance)
	if err != nil {
		return nil
	}

	return instance
}

func getReaderWithExtendedStructForJsonExample() interface{} {
	integer := 0
	uinteger := uint(0)
	str := ""
	float := 0.0
	boolean := false

	instance := dynamicstruct.ExtendStruct(struct {
		Integer   int     `json:"int" validate:"lt=123"`
		Uinteger  uint    `validate:"gte=0"`
		Text      string  `json:"someText"`
		Float     float64 `json:"double"`
		Boolean   bool
		Slice     []int
		Time      time.Time
		SubStruct struct {
			Integer int
			Text    string `json:"subText"`
		} `json:"subData"`
	}{}).
		AddField("PointerInteger", &integer, "").
		AddField("PointerUinteger", &uinteger, "").
		AddField("PointerText", &str, "").
		AddField("PointerFloat", &float, "").
		AddField("PointerBoolean", &boolean, "").
		AddField("PointerTime", &time.Time{}, "").
		AddField("Anonymous", "", `json:"-" validate:"required"`).
		Build().
		New()

	data := []byte(`
{
	"int": 123,
	"Uinteger": 456,
	"someText": "example",
	"double": 123.45,
	"Boolean": true,
	"Time": "2018-12-27T19:42:31+07:00",
	"Slice": [1, 2, 3],
	"PointerInteger": 345,
	"PointerUinteger": 234,
	"PointerFloat": 567.89,
	"PointerText": "pointer example",
	"PointerBoolean": true,
	"PointerTime": "2018-12-28T01:23:45+07:00",
	"subData": {
		"Integer": 12,
		"subText": "sub example"
	},
	"Anonymous": "avoid to read"
}
`)

	err := json.Unmarshal(data, &instance)
	if err != nil {
		return nil
	}

	return instance
}

func getReaderWithMergedStructsForJsonExample() interface{} {
	instance := dynamicstruct.MergeStructs(struct {
		Integer  int     `json:"int" validate:"lt=123"`
		Uinteger uint    `validate:"gte=0"`
		Text     string  `json:"someText"`
		Float    float64 `json:"double"`
		Boolean  bool
		Slice    []int
		Time     time.Time
	}{}, struct {
		Anonymous string `json:"-" validate:"required"`
		SubStruct struct {
			Integer int
			Text    string `json:"subText"`
		} `json:"subData"`
	}{}, struct {
		PointerInteger  *int
		PointerUinteger *uint
		PointerText     *string
		PointerFloat    *float64
		PointerBoolean  *bool
		PointerTime     *time.Time
	}{}).
		Build().
		New()

	data := []byte(`
{
	"int": 123,
	"Uinteger": 456,
	"someText": "example",
	"double": 123.45,
	"Boolean": true,
	"Time": "2018-12-27T19:42:31+07:00",
	"Slice": [1, 2, 3],
	"PointerInteger": 345,
	"PointerUinteger": 234,
	"PointerFloat": 567.89,
	"PointerText": "pointer example",
	"PointerBoolean": true,
	"PointerTime": "2018-12-28T01:23:45+07:00",
	"subData": {
		"Integer": 12,
		"subText": "sub example"
	},
	"Anonymous": "avoid to read"
}
`)

	err := json.Unmarshal(data, &instance)
	if err != nil {
		return nil
	}

	return instance
}
