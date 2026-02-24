package form

import (
	"fmt"
	"net/url"
	"reflect"

	"codeberg.org/aur0ra/form/field"
)

type Decoder[T any] struct {
	fields []field.Field
}

func (d *Decoder[T]) Decode(values url.Values) (decoded T, err error) {
	dp := reflect.ValueOf(&decoded)
	e := dp.Elem()

	for _, f := range d.fields {
		err = f.Decode(values, e.FieldByName(f.Name))
		if err != nil {
			return decoded, err
		}
	}

	return decoded, nil
}

func NewDecoder[T any]() (*Decoder[T], error) {
	t := reflect.TypeFor[T]()
	if t.Kind() != reflect.Struct {
		return nil, InvalidTypeErr{Kind: t.Kind()}
	}

	var decoder Decoder[T]
	decoder.fields = make([]field.Field, 0, t.NumField())

	for structField := range t.Fields() {
		if !structField.IsExported() {
			continue
		}

		f, err := field.New(structField)
		if err != nil {
			continue
		}

		decoder.fields = append(decoder.fields, f)
	}

	return &decoder, nil
}

func MustNewDecoder[T any]() *Decoder[T] {
	d, err := NewDecoder[T]()
	if err != nil {
		panic(err)
	}
	return d
}

type InvalidTypeErr struct {
	Kind reflect.Kind
}

func (i InvalidTypeErr) Error() string {
	return fmt.Sprintf("invalid type: %s", i.Kind.String())
}
