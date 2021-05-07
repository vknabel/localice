package export

import (
	"fmt"
	"io"
	"strconv"

	"github.com/vknabel/localice/internal"
)

type StringsLocalizationExporter struct {
	w io.Writer
}

func NewStringsLocalizationExporter(w io.Writer) StringsLocalizationExporter {
	return StringsLocalizationExporter{w}
}

func (stringsWriter StringsLocalizationExporter) Export(localization internal.Localization) error {
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
