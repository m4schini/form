package form

import (
	"net/url"
	"reflect"
)

type field struct {
	Decoder     FieldDecoder
	MetaData    fieldTag
	StructField reflect.StructField
	_fieldName  string
}

func newField(f reflect.StructField) (field, error) {
	md := parseFieldTag(f.Tag.Get(structTageName))

	fieldName := f.Name
	if md.Alias != "" {
		fieldName = md.Alias
	}

	decoder, err := decoderFor(f.Type)
	if err != nil {
		return field{}, err
	}

	return field{
		Decoder:     decoder,
		MetaData:    md,
		StructField: f,
		_fieldName:  fieldName,
	}, nil
}

func (f *field) Decode(values url.Values, target reflect.Value) error {
	entries, exists := values[f._fieldName]
	if !exists || len(entries) == 0 {
		if f.MetaData.Required {
			return RequiredFieldMissingErr{Field: f._fieldName}
		}
		return nil
	}
	value := entries[0]

	v, err := f.Decoder.DecodeForm(value)
	if err != nil {
		return err
	}

	target.Set(reflect.ValueOf(v))
	return nil
}

func decoderFor(t reflect.Type) (FieldDecoder, error) {
	if t.Implements(fieldDecoderType) {
		v := reflect.New(t)
		return v.Interface().(FieldDecoder), nil
	}

	k := t.Kind()
	switch k {
	case reflect.String:
		return StringDecoder{}, nil
	case reflect.Int:
		return IntDecoder{}, nil
	case reflect.Uint:
		return UintDecoder{}, nil
	case reflect.Int8:
		return Int8Decoder{}, nil
	case reflect.Uint8:
		return Uint8Decoder{}, nil
	case reflect.Int16:
		return Int16Decoder{}, nil
	case reflect.Uint16:
		return Uint16Decoder{}, nil
	case reflect.Int32:
		return Int32Decoder{}, nil
	case reflect.Uint32:
		return Uint32Decoder{}, nil
	case reflect.Int64:
		return Int64Decoder{}, nil
	case reflect.Uint64:
		return Uint64Decoder{}, nil
	case reflect.Bool:
		return BoolDecoder{}, nil
	case reflect.Float32:
		return Float32Decoder{}, nil
	case reflect.Float64:
		return Float64Decoder{}, nil
	case reflect.Complex64:
		return Complex64Decoder{}, nil
	case reflect.Complex128:
		return Complex128Decoder{}, nil
	case reflect.Array:
		return nil, MissingDecoderErr{Kind: k}
	case reflect.Slice:
		return nil, MissingDecoderErr{Kind: k}
	case reflect.Interface:
		return nil, MissingDecoderErr{Kind: k}
	case reflect.Struct:
		return nil, MissingDecoderErr{Kind: k}
	case reflect.Chan:
		return nil, MissingDecoderErr{Kind: k}
	case reflect.Func:
		return nil, MissingDecoderErr{Kind: k}
	case reflect.Map:
		return nil, MissingDecoderErr{Kind: k}
	case reflect.Pointer:
		return nil, MissingDecoderErr{Kind: k}
	case reflect.UnsafePointer:
		return nil, MissingDecoderErr{Kind: k}
	}

	return nil, MissingDecoderErr{Kind: k}
}
