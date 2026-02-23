form
===

Package aur0ra/form converts form values (`url.Values`) into structs.

## Example
Here's a quick example: we parse POST form values and then decode them into a struct:

### Quickstart

```go
package main

import (
	"fmt"
	"net/url"

	"codeberg.org/aur0ra/form"
)

type User struct {
	ID   uint64 `form:"id,required"`
	Name string `form:"name"`
	Age  int    `form:"age"`
}

func main() {
	formValues := ParseForm()
	user, _ := form.Decode[User](formValues)
	fmt.Println(user)
}

func ParseForm() url.Values {
	return url.Values{
		"id": []string{"161"},
		"name": []string{"Username"},
		"age":  []string{"42"},
	}
}
```

### Performance
Performance can be more than doubles by reusing the decoder:
```go
func main() {
	decoder, _ := form.NewDecoder[User]()
	
	formValues := ParseForm()
	user, _ := decoder.Decode(formValues)
	fmt.Println(user)
}
```

The supported field types in the struct are:
- string
- int variants (int, int8, int16, int32, int64)
- uint variants (uint, uint8, uint16, uint32, uint64)
- bool
- float variants (float32, float64)
- ~~a pointer to one of the above types~~ (Coming Soon)

Custom types can be decoded when they implement `field.FieldDecoder`.

```go
type FieldDecoder interface {
	DecodeForm(field string) (any, error)
}
```
