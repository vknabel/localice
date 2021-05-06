package internal

import (
	"io"
)

type StringsLocalizationWriter struct {
	w io.Writer
}

func NewStringsLocalizationWriter(w io.Writer) StringsLocalizationWriter {
	return StringsLocalizationWriter{w}
}

func (stringsWriter StringsLocalizationWriter) Write(localization Localization) error {
	for _, translation := range localization.Translations {
		_, err := io.WriteString(stringsWriter.w, SerializedString(translation.Key, translation.Text))
		if err != nil {
			return err
		}
	}
	return nil
}
