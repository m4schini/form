package form

import (
	"reflect"
	"strconv"
)

type FieldDecoder interface {
	DecodeForm(field string) (any, error)
}

var fieldDecoderType = reflect.TypeFor[FieldDecoder]()

type UnknownDecoder struct{}

func (UnknownDecoder) DecodeForm(field string) (any, error) {
	return field, nil
}

type StringDecoder struct{}

func (StringDecoder) DecodeForm(field string) (any, error) {
	return field, nil
}

type IntDecoder struct{}

func (IntDecoder) DecodeForm(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 64)
	return int(i), err
}

type UintDecoder struct{}

func (UintDecoder) DecodeForm(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 64)
	return uint(i), err
}

type Int8Decoder struct{}

func (Int8Decoder) DecodeForm(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 8)
	return int8(i), err
}

type Uint8Decoder struct{}

func (Uint8Decoder) DecodeForm(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 8)
	return uint8(i), err
}

type Int16Decoder struct{}

func (Int16Decoder) DecodeForm(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 16)
	return int16(i), err
}

type Uint16Decoder struct{}

func (Uint16Decoder) DecodeForm(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 16)
	return uint16(i), err
}

type Int32Decoder struct{}

func (Int32Decoder) DecodeForm(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 32)
	return int32(i), err
}

type Uint32Decoder struct{}

func (Uint32Decoder) DecodeForm(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 32)
	return uint32(i), err
}

type Int64Decoder struct{}

func (Int64Decoder) DecodeForm(field string) (any, error) {
	i, err := strconv.ParseInt(field, 10, 64)
	return int64(i), err
}

type Uint64Decoder struct{}

func (Uint64Decoder) DecodeForm(field string) (any, error) {
	i, err := strconv.ParseUint(field, 10, 64)
	return uint64(i), err
}

type BoolDecoder struct{}

func (BoolDecoder) DecodeForm(field string) (any, error) {
	b, err := strconv.ParseBool(field)
	return b, err
}

type Float32Decoder struct{}

func (Float32Decoder) DecodeForm(field string) (any, error) {
	f, err := strconv.ParseFloat(field, 32)
	return float32(f), err
}

type Float64Decoder struct{}

func (Float64Decoder) DecodeForm(field string) (any, error) {
	f, err := strconv.ParseFloat(field, 64)
	return float64(f), err
}

type Complex64Decoder struct{}

func (Complex64Decoder) DecodeForm(field string) (any, error) {
	f, err := strconv.ParseComplex(field, 64)
	return complex64(f), err
}

type Complex128Decoder struct{}

func (Complex128Decoder) DecodeForm(field string) (any, error) {
	f, err := strconv.ParseComplex(field, 128)
	return complex64(f), err
}
