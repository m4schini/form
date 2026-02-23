package form

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	type Form struct {
		Name string
		Age  int `form:"age"`
	}
	var f = Form{
		Name: "Test",
		Age:  42,
	}

	v := Encode[Form](f)
	fmt.Println(v)
}
