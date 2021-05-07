package main

import (
	"fmt"
	"log"

	"github.com/vknabel/localice/config"
	"github.com/vknabel/localice/export"
	"github.com/vknabel/localice/internal"
)

func main() {
	projectDir := "."
	config, err := config.LoadInferredProjectConfig(projectDir)
	if err != nil {
		log.Fatal(err)
	}

	localizationsByName := make(map[string]internal.Localization)
	for _, sourceConfig := range config.CsvSources {
		currentLocalizations, err := internal.ReadSource(projectDir, sourceConfig)
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

	fmt.Println("✅ sources loaded")

	for _, localization := range localizationsByName {
		for _, platform := range config.Exports {
			err = export.ExportLocalization(localization, platform)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println("✅ exports finished")
}
