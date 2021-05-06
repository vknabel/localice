package internal

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type LocalizationWriter interface {
	Write(localization Localization) error
}

func WriteLocalizationForPlatform(localization Localization, platform PlatformConfig) error {
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

func expandedPath(localization Localization, platform PlatformConfig) string {
	path := strings.ReplaceAll(platform.Path, "${lowerLocale}", strings.ToLower(localization.Name))
	path = strings.ReplaceAll(path, "${upperLocale}", strings.ToUpper(localization.Name))
	path = strings.ReplaceAll(path, "${locale}", localization.Name)
	return path
}
