package main

import (
	"log"

	"github.com/vknabel/localice/internal"
)

func main() {
	projectDir := "."
	config, err := internal.LoadProjectConfig(projectDir)
	if err != nil {
		log.Fatal(err)
	}

	localizationsByName := make(map[string]internal.Localization)
	for _, sourceConfig := range config.CsvSources {
		currentLocalizations, err := internal.ReadCsvSource(projectDir, sourceConfig)
		if err != nil {
			log.Fatal(err)
		}
		for _, localization := range currentLocalizations {
			if loc, ok := localizationsByName[localization.Name]; ok {
				loc.Translations = append(loc.Translations, localization.Translations...)
				localizationsByName[localization.Name] = loc
			} else {
				localizationsByName[localization.Name] = localization
			}
		}
	}

	for _, localization := range localizationsByName {
		for _, platform := range config.Platforms {
			err = internal.WriteLocalizationForPlatform(localization, platform)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
