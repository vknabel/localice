package internal

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/language"
)

type LocalizationWriter interface {
	Write(localization Localization) error
}

func WriteLocalizationForExport(localization Localization, platform ExportConfig) error {
	targetPath := expandedPath(localization, platform)
	file, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var writer LocalizationWriter
	switch platform.Format {
	case "resource-xml":
		writer = NewResourceXmlLocalizationWriter(file)
	case "strings":
		writer = NewStringsLocalizationWriter(file)
	default:
		return fmt.Errorf("unknown format: %s", platform.Format)
	}

	filteredLocalization := localization
	filteredLocalization.Translations = make([]Translation, 0, len(localization.Translations))
	for _, translation := range localization.Translations {
		if platform.MatchPlatform != "" {
			matched, err := regexp.MatchString(platform.MatchPlatform, translation.Platform)
			if err != nil {
				return err
			}
			if !matched {
				continue
			}
		}

		if platform.MatchKey != "" {
			matched, err := regexp.MatchString(platform.MatchKey, translation.Key)
			if err != nil {
				return err
			}
			if !matched {
				continue
			}
		}

		filteredLocalization.Translations = append(filteredLocalization.Translations, translation)
	}
	err = writer.Write(filteredLocalization)
	if err != nil {
		return err
	}
	return nil
}

func expandedPath(localization Localization, platform ExportConfig) string {
	path := strings.ReplaceAll(platform.Path, "${lowerLocale}", strings.ToLower(localization.Name))
	path = strings.ReplaceAll(path, "${upperLocale}", strings.ToUpper(localization.Name))
	path = strings.ReplaceAll(path, "${locale}", localization.Name)

	if tag, err := language.Parse(localization.Name); err == nil {
		path = strings.ReplaceAll(path, "${region}", "")

		base, _ := tag.Base()
		path = strings.ReplaceAll(path, "${base}", base.String())
		path = strings.ReplaceAll(path, "${lowerBase}", strings.ToLower(base.String()))
		path = strings.ReplaceAll(path, "${upperBase}", strings.ToUpper(base.String()))

		script, _ := tag.Script()
		path = strings.ReplaceAll(path, "${script}", script.String())
		path = strings.ReplaceAll(path, "${lowerScript}", strings.ToLower(script.String()))
		path = strings.ReplaceAll(path, "${upperScript}", strings.ToUpper(script.String()))

		region, _ := tag.Region()
		path = strings.ReplaceAll(path, "${region}", region.String())
		path = strings.ReplaceAll(path, "${lowerRegion}", strings.ToLower(region.String()))
		path = strings.ReplaceAll(path, "${upperRegion}", strings.ToUpper(region.String()))

	}
	return path
}
