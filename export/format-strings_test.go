package export

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/vknabel/localice/internal"
)

func TestWriteEmptyStringsLocalizationSucceeds(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewStringsLocalizationExporter(buf)
	err := w.Export(internal.Localization{
		Name:         "en",
		Translations: []internal.Translation{},
	})
	if err != nil {
		t.Error(err)
	}
	localizedStrings := buf.String()
	if localizedStrings != "" {
		t.Errorf("not empty: %s", localizedStrings)
	}
}

func TestWriteStringsLocalization(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewStringsLocalizationExporter(buf)
	err := w.Export(internal.Localization{
		Name: "en",
		Translations: []internal.Translation{
			{Platform: "iOS", Key: "Hello %@", Text: "Hello %@!"},
			{Platform: "", Key: "YouShallPass", Text: "You shall pass!"},
			{Platform: "", Key: "Newline", Text: "\n"},
			{Platform: "", Key: "Quotes %@", Text: `"%@"`},
		},
	})
	if err != nil {
		t.Error(err)
	}
	localizedStrings := buf.String()
	expected := `"Hello %@" = "Hello %@!";
"YouShallPass" = "You shall pass!";
"Newline" = "\n";
"Quotes %@" = "\"%@\"";
`
	if localizedStrings != expected {
		t.Errorf("generated wrong contents: %s", localizedStrings)
	}
}

type errorWriter struct {
	err error
}

func (w *errorWriter) Write(p []byte) (n int, err error) {
	return 0, w.err
}

func TestWriteStringsLocalizationPassesErrors(t *testing.T) {
	expectedError := fmt.Errorf("totally expected error")
	w := NewStringsLocalizationExporter(&errorWriter{expectedError})
	err := w.Export(internal.Localization{
		Name: "en",
		Translations: []internal.Translation{
			{Platform: "iOS", Key: "Hello %@", Text: "Hello %@!"},
		},
	})
	if err != expectedError {
		t.Error(err)
	}
}
