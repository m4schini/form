package field

import (
	"reflect"
	"strconv"
)

// Parser converts a string into a type.
// Which type is controlled by the implementation of the Parser.
// Parser must always return the same type.
type Parser interface {
	// ParseForm converts a string into a type.
	ParseForm(field string) (any, error)
}

var fieldDecoderType = reflect.TypeFor[Parser]()

// StringParser returns the field string unchanged
type StringParser struct{}

func (StringParser) ParseForm(field string) (any, error) {
	return field, nil
}

// IntParser parses the field string as a 64 bit int and casts it to int
type IntParser struct{}

func (IntParser) ParseForm(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 64)
	return int(i), err
}

// UintParser parses the field string as a 64 bit uint and casts it to uint
type UintParser struct{}

func (UintParser) ParseForm(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 64)
	return uint(i), err
}

type Int8Parser struct{}

func (Int8Parser) ParseForm(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 8)
	return int8(i), err
}

type Uint8Parser struct{}

func (Uint8Parser) ParseForm(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 8)
	return uint8(i), err
}

type Int16Parser struct{}

func (Int16Parser) ParseForm(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 16)
	return int16(i), err
}

type Uint16Parser struct{}

func (Uint16Parser) ParseForm(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 16)
	return uint16(i), err
}

type Int32Parser struct{}

func (Int32Parser) ParseForm(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 32)
	return int32(i), err
}

type Uint32Parser struct{}

func (Uint32Parser) ParseForm(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 32)
	return uint32(i), err
}

type Int64Parser struct{}

func (Int64Parser) ParseForm(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 64)
	return int64(i), err
}

type Uint64Parser struct{}

func (Uint64Parser) ParseForm(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 64)
	return uint64(i), err
}

type BoolParser struct{}

func (BoolParser) ParseForm(field string) (any, error) {
	b, err := strconv.ParseBool(field)
	return b, err
}

type Float32Parser struct{}

func (Float32Parser) ParseForm(field string) (any, error) {
	f, err := strconv.ParseFloat(field, 32)
	return float32(f), err
}

type Float64Parser struct{}

func (Float64Parser) ParseForm(field string) (any, error) {
	f, err := strconv.ParseFloat(field, 64)
	return float64(f), err
}

type Complex64Parser struct{}

func (Complex64Parser) ParseForm(field string) (any, error) {
	f, err := strconv.ParseComplex(field, 64)
	return complex64(f), err
}

type Complex128Parser struct{}

func (Complex128Parser) ParseForm(field string) (any, error) {
	f, err := strconv.ParseComplex(field, 128)
	return complex64(f), err
}
