form
===

Package form converts form values (`url.Values`) into structs.

## Example
Here's a quick example: we parse POST form values and then decode them into a struct:

### Quickstart

```go
package main

import (
    "fmt"
    "net/http"
    
    "codeberg.org/aur0ra/form"
)

type User struct {
    ID   uint64 `form:"id,required"`
    Name string `form:"name"`
    Age  int    `form:"age"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    
    user, _ := form.Decode[User](r.Form)
	
    fmt.Println(user)
}
```

### Performance
Performance can be more than doubles by reusing the decoder:
```go
var decoder, _ := form.NewDecoder[User]()

func Handler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
	
    user, _ := decoder.Decode(r.Form)
	
    fmt.Println(user)
}
```

The supported field types in the struct are:
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64
- bool
- float32, float64
- complex64, complex128
- ~~a pointer to one of the above types~~ (Coming Soon)

Custom types can be decoded when they implement `field.Parser`.

```go
type Parser interface {
	Parse(field string) (any, error)
}
```
Parser must return the custom type or nil.