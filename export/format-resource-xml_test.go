package export

import (
	"bytes"
	"testing"

	"github.com/vknabel/localice/internal"
)

func TestResourceXmlExportSucceedsWhenEmpty(t *testing.T) {
	expected := `<?xml version="1.0" encoding="utf-8"?>
<resources></resources>
`
	actual, err := exported(internal.Localization{
		Name:         "en",
		Translations: []internal.Translation{},
	})
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("not expected: %s", actual)
	}
}

func TestResourceXmlExportSucceedsWhenValid(t *testing.T) {
	expected := `<?xml version="1.0" encoding="utf-8"?>
<resources>
    <string name="hello_world">Hello World!</string>
</resources>
`
	actual, err := exported(internal.Localization{
		Name: "en",
		Translations: []internal.Translation{
			{Key: "hello_world", Text: "Hello World!"},
		},
	})
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("not expected: %s", actual)
	}
}

func TestResourceXmlExportSucceedsWithSpecialCharacters(t *testing.T) {
	expected := `<?xml version="1.0" encoding="utf-8"?>
<resources>
    <string name="simple_emoji">🙈</string>
    <string name="xml_characters">&lt;xml&gt;</string>
    <string name="apostrophe">\&#39;</string>
</resources>
`
	actual, err := exported(internal.Localization{
		Name: "en",
		Translations: []internal.Translation{
			{Key: "simple_emoji", Text: "🙈"},
			{Key: "xml_characters", Text: "<xml>"},
			{Key: "apostrophe", Text: "'"},
		},
	})
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("not expected: \n%s\n\nexpected: \n%s", actual, expected)
	}
}

func exported(localization internal.Localization) (string, error) {
	buf := new(bytes.Buffer)
	w := NewResourceXmlLocalizationExporter(buf)
	err := w.Export(localization)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
