package internal

import (
	"fmt"
	"io"
	"strconv"
)

type StringsLocalizationWriter struct {
	w io.Writer
}

func NewStringsLocalizationWriter(w io.Writer) StringsLocalizationWriter {
	return StringsLocalizationWriter{w}
}

func (stringsWriter StringsLocalizationWriter) Write(localization Localization) error {
	for _, translation := range localization.Translations {
		_, err := io.WriteString(stringsWriter.w, serializedString(translation.Key, translation.Text))
		if err != nil {
			return err
		}
	}
	return nil
}

func serializedString(key string, value string) string {
	return fmt.Sprintf("%s = %s;\n", quoted(key), quoted(value))
}

func quoted(value string) string {
	return strconv.Quote(value)
}
