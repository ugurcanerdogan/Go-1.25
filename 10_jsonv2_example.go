package main

import (
	"bytes"
	"encoding/json/jsontext"
	jsonv2 "encoding/json/v2"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var p Person

func main() {
	// IDE Settings: experimental=jsonv2
	//
	JoinCustomMarshalers()
	//
	alice, err := SimpleUnmarshal()
	Options(alice)
	//
	StreamDecode()
	MarshalWrite()
	UnmarshalRead()
	//
	ErrorLogEnhancement(err)
}

func Options(person Person) {
	data, err := jsonv2.Marshal(
		person,
	)
	fmt.Println(string(data), err)

	data2, err := jsonv2.Marshal(
		person,
		jsonv2.OmitZeroStructFields(true),
		jsonv2.JoinOptions(
			jsontext.SpaceAfterColon(true),
			jsontext.SpaceAfterComma(true),
		),
	)
	fmt.Println(string(data2), err)
}

func StreamDecode() {
	stream := `[{"name":"Alice","age":30},{"name":"Bob","age":25}]`
	dec := jsontext.NewDecoder(bytes.NewReader([]byte(stream)))
	for {
		var p Person
		if err := jsonv2.UnmarshalDecode(dec, &p); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("decode error:", err)
			break
		}
		fmt.Println(p)
	}
}

func MarshalWrite() {
	alice := Person{"Alice", 30}
	jsonv2.MarshalWrite(os.Stdout, alice)
	fmt.Println()
}

func UnmarshalRead() {
	in := bytes.NewReader([]byte(`{"name":"Bob","age":25}`))
	var bob Person
	jsonv2.UnmarshalRead(in, &bob)
	fmt.Println(bob)
}

func JoinCustomMarshalers() {
	// Custom marshaler simple example
	boolMarshaler := jsonv2.MarshalToFunc(
		func(enc *jsontext.Encoder, val bool) error {
			if val {
				return enc.WriteToken(jsontext.String("✓"))
			}
			return enc.WriteToken(jsontext.String("✗"))
		},
	)
	strMarshaler := jsonv2.MarshalToFunc(
		func(enc *jsontext.Encoder, val string) error {
			if val == "on" || val == "true" {
				return enc.WriteToken(jsontext.String("✓"))
			}
			if val == "off" || val == "false" {
				return enc.WriteToken(jsontext.String("✗"))
			}
			return jsonv2.SkipFunc
		},
	)
	marshalers := jsonv2.JoinMarshalers(boolMarshaler, strMarshaler)
	vals := []any{true, "off", "hello"}
	data, err := jsonv2.Marshal(vals, jsonv2.WithMarshalers(marshalers))
	fmt.Println(string(data), err)
}

func SimpleUnmarshal() (Person, error) {
	err := jsonv2.Unmarshal([]byte(`{"name":"Alice","age":30}`), &p)
	fmt.Println(p, err)
	return p, err
}

func ErrorLogEnhancement(err error) {
	// better error logs
	fmt.Println("\nError example:")
	err = jsonv2.Unmarshal([]byte(`{"port":"oops"}`), &struct{ Port int }{})
	fmt.Println("error:", err)
}
