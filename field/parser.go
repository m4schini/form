package field

import (
	"reflect"
	"strconv"
)

// Parser converts a string into a type. Which type is controlled by the implementation of the Parser.
type Parser interface {
	Parse(field string) (any, error)
}

var fieldDecoderType = reflect.TypeFor[Parser]()

// StringDecoder returns the field string unchanged
type StringDecoder struct{}

func (StringDecoder) Parse(field string) (any, error) {
	return field, nil
}

// IntDecoder parses the field string as a 64 bit int and casts it to int
type IntDecoder struct{}

func (IntDecoder) Parse(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 64)
	return int(i), err
}

// UintDecoder parses the field string as a 64 bit uint and casts it to uint
type UintDecoder struct{}

func (UintDecoder) Parse(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 64)
	return uint(i), err
}

type Int8Decoder struct{}

func (Int8Decoder) Parse(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 8)
	return int8(i), err
}

type Uint8Decoder struct{}

func (Uint8Decoder) Parse(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 8)
	return uint8(i), err
}

type Int16Decoder struct{}

func (Int16Decoder) Parse(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 16)
	return int16(i), err
}

type Uint16Decoder struct{}

func (Uint16Decoder) Parse(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 16)
	return uint16(i), err
}

type Int32Decoder struct{}

func (Int32Decoder) Parse(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 32)
	return int32(i), err
}

type Uint32Decoder struct{}

func (Uint32Decoder) Parse(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 32)
	return uint32(i), err
}

type Int64Decoder struct{}

func (Int64Decoder) Parse(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 64)
	return int64(i), err
}

type Uint64Decoder struct{}

func (Uint64Decoder) Parse(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 64)
	return uint64(i), err
}

type BoolDecoder struct{}

func (BoolDecoder) Parse(field string) (any, error) {
	b, err := strconv.ParseBool(field)
	return b, err
}

type Float32Decoder struct{}

func (Float32Decoder) Parse(field string) (any, error) {
	f, err := strconv.ParseFloat(field, 32)
	return float32(f), err
}

type Float64Decoder struct{}

func (Float64Decoder) Parse(field string) (any, error) {
	f, err := strconv.ParseFloat(field, 64)
	return float64(f), err
}

type Complex64Decoder struct{}

func (Complex64Decoder) Parse(field string) (any, error) {
	f, err := strconv.ParseComplex(field, 64)
	return complex64(f), err
}

type Complex128Decoder struct{}

func (Complex128Decoder) Parse(field string) (any, error) {
	f, err := strconv.ParseComplex(field, 128)
	return complex64(f), err
}
