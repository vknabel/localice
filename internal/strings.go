package internal

import (
	"fmt"
	"strconv"
)

func SerializedString(key string, value string) string {
	return fmt.Sprintf("%s = %s;\n", quoted(key), quoted(value))
}

func quoted(value string) string {
	return strconv.Quote(value)
}
