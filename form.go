package form

import (
	"fmt"
	"net/url"
	"reflect"

	field "codeberg.org/aur0ra/form/field"
)

func Decode[T any](values url.Values) (decoded T, err error) {
	d, err := NewDecoder[T]()
	if err != nil {
		return decoded, err
	}

	return d.Decode(values)
}

func Encode[T any](form T) url.Values {
	t := reflect.TypeFor[T]()
	v := reflect.ValueOf(form)
	values := make(url.Values)
	for structField := range t.Fields() {
		tag := field.ParseTag(structField.Tag.Get(field.TagName))
		name := tag.FieldName(structField.Name)
		f := v.FieldByName(name)
		//str := f.Interface().(fmt.Stringer)
		values.Set(name, fmt.Sprint(f))
	}

	return values
}
