# Golang dynamic struct

Package dynamic struct provides possibility to dynamically, in runtime,
extend or merge existing defined structs or to provide completely new struct.

Main features:
* Building completely new struct in runtime
* Extending existing struct in runtime
* Merging multiple structs in runtime
* Adding new fields into struct
* Removing existing fields from struct
* Modifying fields' types and tags
* Easy reading of dynamic structs

Works out-of-the-box with:
* https://github.com/go-playground/form
* https://github.com/go-playground/validator
* https://github.com/leebenson/conform
* https://golang.org/pkg/encoding/json/
* ...

## Benchmarks

Environment:
* MacBook Pro (13-inch, Early 2015), 2,7 GHz Intel Core i5
* go version go1.11 darwin/amd64

```
goos: darwin
goarch: amd64
pkg: github.com/ompluscator/dynamic-struct
BenchmarkClassicWay_NewInstance-4                 2000000000     0.34 ns/op
BenchmarkNewStruct_NewInstance-4                    10000000      141 ns/op
BenchmarkNewStruct_NewInstance_Parallel-4           20000000     89.6 ns/op
BenchmarkExtendStruct_NewInstance-4                 10000000      135 ns/op
BenchmarkExtendStruct_NewInstance_Parallel-4        20000000     89.5 ns/op
BenchmarkMergeStructs_NewInstance-4                 10000000      140 ns/op
BenchmarkMergeStructs_NewInstance_Parallel-4        20000000     94.3 ns/op
```

## Add new struct
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ompluscator/dynamic-struct"
)

func main() {
	instance := dynamicstruct.NewStruct().
		AddField("Integer", 0, `json:"int"`).
		AddField("Text", "", `json:"someText"`).
		AddField("Float", 0.0, `json:"double"`).
		AddField("Boolean", false, "").
		AddField("Slice", []int{}, "").
		AddField("Anonymous", "", `json:"-"`).
		Build().
		New()

	data := []byte(`
{
    "int": 123,
    "someText": "example",
    "double": 123.45,
    "Boolean": true,
    "Slice": [1, 2, 3],
    "Anonymous": "avoid to read"
}
`)

	err := json.Unmarshal(data, &instance)
	if err != nil {
		log.Fatal(err)
	}

	data, err = json.Marshal(instance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
	// Out:
	// {"int":123,"someText":"example","double":123.45,"Boolean":true,"Slice":[1,2,3]}
}
```

## Extend existing struct
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ompluscator/dynamic-struct"
)

type Data struct {
	Integer int `json:"int"`
}

func main() {
	instance := dynamicstruct.ExtendStruct(Data{}).
		AddField("Text", "", `json:"someText"`).
		AddField("Float", 0.0, `json:"double"`).
		AddField("Boolean", false, "").
		AddField("Slice", []int{}, "").
		AddField("Anonymous", "", `json:"-"`).
		Build().
		New()

	data := []byte(`
{
    "int": 123,
    "someText": "example",
    "double": 123.45,
    "Boolean": true,
    "Slice": [1, 2, 3],
    "Anonymous": "avoid to read"
}
`)

	err := json.Unmarshal(data, &instance)
	if err != nil {
		log.Fatal(err)
	}

	data, err = json.Marshal(instance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
	// Out:
	// {"int":123,"someText":"example","double":123.45,"Boolean":true,"Slice":[1,2,3]}
}
```

## Merge existing structs
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ompluscator/dynamic-struct"
)

type DataOne struct {
	Integer int     `json:"int"`
	Text    string  `json:"someText"`
	Float   float64 `json:"double"`
}

type DataTwo struct {
	Boolean bool
	Slice []int
	Anonymous string `json:"-"`
}

func main() {
	instance := dynamicstruct.MergeStructs(DataOne{}, DataTwo{}).
		Build().
		New()

	data := []byte(`
{
"int": 123,
"someText": "example",
"double": 123.45,
"Boolean": true,
"Slice": [1, 2, 3],
"Anonymous": "avoid to read"
}
`)

	err := json.Unmarshal(data, &instance)
	if err != nil {
		log.Fatal(err)
	}

	data, err = json.Marshal(instance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
	// Out:
	// {"int":123,"someText":"example","double":123.45,"Boolean":true,"Slice":[1,2,3]}
}
```

## Read dynamic struct

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ompluscator/dynamic-struct"
)

func main() {
	instance := dynamicstruct.NewStruct().
		AddField("Integer", 0, `json:"int"`).
		AddField("Text", "", `json:"someText"`).
		AddField("Float", 0.0, `json:"double"`).
		AddField("Boolean", false, "").
		AddField("Slice", []int{}, "").
		AddField("Anonymous", "", `json:"-"`).
		Build().
		New()

	data := []byte(`
{
    "int": 123,
    "someText": "example",
    "double": 123.45,
    "Boolean": true,
    "Slice": [1, 2, 3],
    "Anonymous": "avoid to read"
}
`)

	err := json.Unmarshal(data, &instance)
	if err != nil {
		log.Fatal(err)
	}

	value := dynamicstruct.NewReader(instance)
	fmt.Println("Integer", value.GetField("Integer").Int())
	fmt.Println("Text", value.GetField("Text").String())
	fmt.Println("Float", value.GetField("Float").Float64())
	fmt.Println("Boolean", value.GetField("Boolean").Bool())
	fmt.Println("Slice", value.GetField("Slice").Interface().([]int))
	fmt.Println("Anonymous", value.GetField("Anonymous").String())

	// Out:
	// Integer 123
	// Text example
	// Float 123.45
	// Boolean true
	// Slice [1 2 3]
	// Anonymous
}
```