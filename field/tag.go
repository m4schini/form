package field

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	TagName = "form"
)

// Tag represents the configuration options that can be with the struct tag (see TagName) set on struct fields.
type Tag struct {
	// If Alias is not an empty string, it will be used instead of the struct field name.
	Alias string
	// If Required is true decode will fail if FieldName is not present in url.Values.
	Required bool
	// If Regex is not nil it will be used to validate the field value before decoding it.
	Regex *regexp.Regexp
}

// FieldName returns the key for the url.Values lookup.
func (t Tag) FieldName(fieldName string) string {
	if t.Alias != "" {
		fieldName = t.Alias
	}
	return fieldName
}

// ParseTag parses the string of the struct tag into a Tag struct
func ParseTag(tag string) Tag {
	if tag == "" {
		return Tag{}
	}
	parts := strings.Split(tag, ",")
	if len(parts) == 0 {
		return Tag{}
	}
	if len(parts) == 1 {
		return Tag{Alias: parts[0]}
	}
	if len(parts) == 2 {
		aliasPart := parts[0]
		requiredPart := parts[1]

		required := requiredPart == "required"
		if !required {
			requiredParts := strings.Split(requiredPart, "=")
			if len(requiredParts) == 2 {
				required, _ = strconv.ParseBool(requiredParts[1])
			}
		}

		return Tag{Alias: aliasPart, Required: required}
	}

	return Tag{}
}
