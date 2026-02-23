package field

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	TagName = "form"
)

type Tag struct {
	Alias    string
	Required bool
	Regex    *regexp.Regexp
}

func (t Tag) FieldName(fieldName string) string {
	if t.Alias != "" {
		fieldName = t.Alias
	}
	return fieldName
}

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
