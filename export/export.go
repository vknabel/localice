package export

import (
	"fmt"
	"os"
	"regexp"

	"github.com/vknabel/localice/config"
	"github.com/vknabel/localice/internal"
)

type LocalizationExporter interface {
	Export(localization internal.Localization) error
}

func ExportLocalization(localization internal.Localization, export config.ExportConfig) error {
	targetPath := expandExportPath(localization.Name, export.Path)
	file, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var writer LocalizationExporter
	switch export.Format {
	case "resource-xml":
		writer = NewResourceXmlLocalizationExporter(file)
	case "strings":
		writer = NewStringsLocalizationExporter(file)
	default:
		return fmt.Errorf("unknown format: %s", export.Format)
	}

	filteredLocalization, err := matchingLocalizationForExport(localization, export)
	if err != nil {
		return err
	}
	err = writer.Export(filteredLocalization)
	if err != nil {
		return err
	}
	return nil
}

func matchingLocalizationForExport(localization internal.Localization, export config.ExportConfig) (internal.Localization, error) {
	filteredLocalization := localization
	filteredLocalization.Translations = make([]internal.Translation, 0, len(localization.Translations))
	for _, translation := range localization.Translations {
		if export.MatchPlatform != "" {
			matched, err := regexp.MatchString(export.MatchPlatform, translation.Platform)
			if err != nil {
				return localization, err
			}
			if !matched {
				continue
			}
		}

		if export.MatchKey != "" {
			matched, err := regexp.MatchString(export.MatchKey, translation.Key)
			if err != nil {
				return localization, err
			}
			if !matched {
				continue
			}
		}

		filteredLocalization.Translations = append(filteredLocalization.Translations, translation)
	}
	return filteredLocalization, nil
}
