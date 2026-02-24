package form

import (
	"errors"
	"net/url"
	"testing"
	"time"

	"codeberg.org/aur0ra/form/field"
)

type Form struct {
	Name string
	Days int
}

func TestDecode_BuiltinTypes(t *testing.T) {
	var expected = Form{
		Name: "Test Name",
		Days: 42,
	}
	var raw = url.Values{
		"Name": []string{"Test Name"},
		"Days": []string{"42"},
	}

	actual, err := Decode[Form](raw)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(actual)

	if actual.Name != expected.Name {
		t.Log("name does not match:", actual.Name, expected.Name)
		t.Fail()
	}
	if actual.Days != expected.Days {
		t.Log("days does not match:", actual.Days, expected.Days)
		t.Fail()
	}
}

func TestDecode_Ptr(t *testing.T) {
	var raw = url.Values{}
	_, err := Decode[*Form](raw)
	t.Log(err)
	if _, correctErr := errors.AsType[InvalidTypeErr](err); !correctErr {
		t.Fail()
	}
}

type Date time.Time

func (d Date) String() string {
	return time.Time(d).String()
}

func (d Date) Parse(field string) (any, error) {
	t, err := time.Parse("2006-01-02", field)
	return Date(t), err
}

type FormWithTime struct {
	Name string
	Days int
	Date Date
}

func TestDecode_CustomType(t *testing.T) {
	dv, _ := time.Parse("2006-01-02", "2020-06-01")
	var expected = FormWithTime{
		Name: "Test Name",
		Days: 42,
		Date: Date(dv),
	}
	var raw = url.Values{
		"Name": []string{"Test Name"},
		"Days": []string{"42"},
		"Date": []string{"2020-06-01"},
	}

	actual, err := Decode[FormWithTime](raw)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(actual)

	if actual.Name != expected.Name {
		t.Log("name does not match:", actual.Name, expected.Name)
		t.Fail()
	}
	if actual.Days != expected.Days {
		t.Log("days does not match:", actual.Days, expected.Days)
		t.Fail()
	}
	if actual.Date != expected.Date {
		t.Log("date does not match:", actual.Date, expected.Date)
		t.Fail()
	}
}

func TestDecode_Alias(t *testing.T) {
	type FormWithAlias struct {
		Name string `form:"alias"`
	}
	var expected = FormWithAlias{
		Name: "Test Name",
	}
	var raw = url.Values{
		"alias": []string{"Test Name"},
	}

	actual, err := Decode[FormWithAlias](raw)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(actual)

	if actual.Name != expected.Name {
		t.Log("name does not match:", actual.Name, expected.Name)
		t.Fail()
	}
}

func TestDecode_Required(t *testing.T) {
	type FormWithAlias struct {
		Name string `form:"alias,required"`
	}
	var raw = url.Values{
		"Required": []string{"Test Name"},
	}

	actual, err := Decode[FormWithAlias](raw)
	if _, expectedErr := errors.AsType[field.RequiredFieldMissingErr](err); !expectedErr {
		t.Fail()
	}
	t.Log(err, actual)
}

func TestDecode_RequiredKV(t *testing.T) {
	type FormWithAlias struct {
		Name string `form:"alias,required=true"`
	}
	var raw = url.Values{
		"Required": []string{"Test Name"},
	}

	actual, err := Decode[FormWithAlias](raw)
	if _, expectedErr := errors.AsType[field.RequiredFieldMissingErr](err); !expectedErr {
		t.Fail()
	}
	t.Log(err, actual)
}

type UnknownType int

type BenchmarkForm struct {
	Name  string
	Days  int
	Index uint8
	Date  Date
	F64   float64 `form:"form64"`
	Unk   UnknownType
	Strs  []string
	Nums  []int
}

var rawBenchmarkForm = url.Values{
	"Name":   []string{"Test Name"},
	"Days":   []string{"42"},
	"Index":  []string{"0"},
	"Date":   []string{"2020-06-01"},
	"form64": []string{"42.69"},
	"Strs":   []string{"one,two,three"},
	"Nums":   []string{"1,2,3"},
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Decode[BenchmarkForm](rawBenchmarkForm)
		if err != nil {
			b.Log(err)
			b.Fail()
		}
	}
}

func BenchmarkDecoder(b *testing.B) {
	var decoder, err = NewDecoder[BenchmarkForm]()
	if err != nil {
		b.Log(err)
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		_, err := decoder.Decode(rawBenchmarkForm)
		if err != nil {
			b.Log(err)
			b.Fail()
		}
	}
}
