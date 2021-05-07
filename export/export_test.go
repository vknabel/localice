package export

import (
	"testing"

	"github.com/vknabel/localice/config"
	"github.com/vknabel/localice/internal"
)

func TestMatchingLocalizationForExportEqualsWhenAllMatch(t *testing.T) {
	unfiltered := internal.Localization{
		Name: "en",
		Translations: []internal.Translation{
			{Key: "hello_world", Text: "Hello World!"},
			{Key: "key", Text: "It's a key!"},
		},
	}
	exportConfig := config.ExportConfig{}
	filtered, err := matchingLocalizationForExport(unfiltered, exportConfig)
	if err != nil {
		t.Error(err)
	}
	if filtered.Name != unfiltered.Name {
		t.Errorf("name must not change: %s", filtered)
	}
	if len(unfiltered.Translations) != len(filtered.Translations) {
		t.Errorf("must have same length: %d", len(filtered.Translations))
	}
	for index := range unfiltered.Translations {
		if filtered.Translations[index] != unfiltered.Translations[index] {
			t.Errorf("order and translation must equal: %d %s", index, filtered)
		}
	}
}

func TestMatchingLocalizationForExportFiltersNonMatching(t *testing.T) {
	unfiltered := internal.Localization{
		Name: "en",
		Translations: []internal.Translation{
			{Platform: "iOS", Key: "hello_world", Text: "Hello World!"},
			{Platform: "", Key: "empty", Text: "Empty"},
			{Platform: "iOS", Key: "NSCameraUsageDescription", Text: "To take pictures and so"},
			{Platform: "Android", Key: "key", Text: "It's a key!"},
		},
	}
	exportConfig := config.ExportConfig{
		MatchPlatform: "^(iOS)?$",
		MatchKey:      "^([^N][^S])",
	}
	filtered, err := matchingLocalizationForExport(unfiltered, exportConfig)
	if err != nil {
		t.Error(err)
	}
	if filtered.Name != unfiltered.Name {
		t.Errorf("name must not change: %s", filtered)
	}
	expected := internal.Localization{
		Name: "en",
		Translations: []internal.Translation{
			{Platform: "iOS", Key: "hello_world", Text: "Hello World!"},
			{Platform: "", Key: "empty", Text: "Empty"},
		},
	}
	if len(expected.Translations) != len(filtered.Translations) {
		t.Errorf("must have same length: expected %d actual %d", len(expected.Translations), len(filtered.Translations))
	}
	for index := range expected.Translations {
		if filtered.Translations[index] != expected.Translations[index] {
			t.Errorf("unexpected filter result: %d %s", index, filtered)
		}
	}
}

func TestMatchingLocalizationFailsForInvalidMatchPlatform(t *testing.T) {
	unfiltered := internal.Localization{
		Name: "en",
		Translations: []internal.Translation{
			{Platform: "Platform", Key: "hello_world", Text: "Hello World!"},
			{Platform: "WillBeIgnored", Key: "key", Text: "It's a key!"},
		},
	}
	_, err := matchingLocalizationForExport(unfiltered, config.ExportConfig{
		MatchPlatform: "\\",
	})
	if err == nil {
		t.Error("expected error for MatchPlatform")
	}
	_, err = matchingLocalizationForExport(unfiltered, config.ExportConfig{
		MatchKey: "\\",
	})
	if err == nil {
		t.Error("expected error for MatchKey")
	}
}

func TestExportLocalizationSucceedsWhenEmpty(t *testing.T) {

}
