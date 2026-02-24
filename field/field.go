package field

import (
	"fmt"
	"net/url"
	"reflect"
)

// Field contains the metadata and type parser to decode a key-value pair from url.Values into a struct field.
type Field struct {
	// Name is the name of the field inside the struct
	Name string
	// Parser used to turn string value into Field type
	Parser Parser
	// MetaData contains additional Tag metadata used during decoding
	MetaData Tag
	// _fieldName is the key used during Decode to get the value from url.Values
	_fieldName string
}

func New(f reflect.StructField) (Field, error) {
	md := ParseTag(f.Tag.Get(TagName))

	fieldName := md.FieldName(f.Name)

	parser, err := parserFor(f.Type)
	if err != nil {
		return Field{}, err
	}

	return Field{
		Name:       f.Name,
		Parser:     parser,
		MetaData:   md,
		_fieldName: fieldName,
	}, nil
}

// Decode retrieves value from url.Values, converts its value into the target type and then sets the value on target
func (f *Field) Decode(values url.Values, target reflect.Value) error {
	entries, exists := values[f._fieldName]
	if !exists || len(entries) == 0 {
		if f.MetaData.Required {
			return RequiredFieldMissingErr{Field: f._fieldName}
		}
		return nil
	}
	value := entries[0]

	v, err := f.Parser.Parse(value)
	if err != nil {
		return err
	}

	target.Set(reflect.ValueOf(v))
	return nil
}

// parserFor is a factory method that returns the correct parser to convert string into type t.
func parserFor(t reflect.Type) (Parser, error) {
	if t.Implements(fieldDecoderType) {
		v := reflect.New(t)
		return v.Interface().(Parser), nil
	}

	k := t.Kind()
	switch k {
	case reflect.String:
		return StringParser{}, nil
	case reflect.Int:
		return IntParser{}, nil
	case reflect.Uint:
		return UintParser{}, nil
	case reflect.Int8:
		return Int8Parser{}, nil
	case reflect.Uint8:
		return Uint8Parser{}, nil
	case reflect.Int16:
		return Int16Parser{}, nil
	case reflect.Uint16:
		return Uint16Parser{}, nil
	case reflect.Int32:
		return Int32Parser{}, nil
	case reflect.Uint32:
		return Uint32Parser{}, nil
	case reflect.Int64:
		return Int64Parser{}, nil
	case reflect.Uint64:
		return Uint64Parser{}, nil
	case reflect.Bool:
		return BoolParser{}, nil
	case reflect.Float32:
		return Float32Parser{}, nil
	case reflect.Float64:
		return Float64Parser{}, nil
	case reflect.Complex64:
		return Complex64Parser{}, nil
	case reflect.Complex128:
		return Complex128Parser{}, nil
	case reflect.Array:
		return nil, CannotParseErr{Kind: k}
	case reflect.Slice:
		return nil, CannotParseErr{Kind: k}
	case reflect.Interface:
		return nil, CannotParseErr{Kind: k}
	case reflect.Struct:
		return nil, CannotParseErr{Kind: k}
	case reflect.Chan:
		return nil, CannotParseErr{Kind: k}
	case reflect.Func:
		return nil, CannotParseErr{Kind: k}
	case reflect.Map:
		return nil, CannotParseErr{Kind: k}
	case reflect.Pointer:
		return nil, CannotParseErr{Kind: k}
	case reflect.UnsafePointer:
		return nil, CannotParseErr{Kind: k}
	case reflect.Invalid:
		return nil, CannotParseErr{Kind: k}
	case reflect.Uintptr:
		return nil, CannotParseErr{Kind: k}
	}

	return nil, CannotParseErr{Kind: k}
}

type RequiredFieldMissingErr struct {
	Field string
}

func (r RequiredFieldMissingErr) Error() string {
	return fmt.Sprintf("required field is missing: %v", r.Field)
}

type CannotParseErr struct {
	Kind reflect.Kind
}

func (i CannotParseErr) Error() string {
	return fmt.Sprintf("type has no parser: %s", i.Kind.String())
}
