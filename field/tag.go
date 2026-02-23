package form

import (
	"strconv"
	"strings"
)

type fieldTag struct {
	Alias    string
	Required bool
	Regex    string
}

func parseFieldTag(tag string) fieldTag {
	if tag == "" {
		return fieldTag{}
	}
	parts := strings.Split(tag, ",")
	if len(parts) == 0 {
		return fieldTag{}
	}
	if len(parts) == 1 {
		return fieldTag{Alias: parts[0]}
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

		return fieldTag{Alias: aliasPart, Required: required}
	}

	return fieldTag{}
}
